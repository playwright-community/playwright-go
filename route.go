package playwright

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type routeImpl struct {
	channelOwner
	handling *chan bool
}

func (r *routeImpl) startHandling() chan bool {
	r.Lock()
	defer r.Unlock()
	handling := make(chan bool, 1)
	r.handling = &handling
	return *r.handling
}

func (r *routeImpl) reportHandled(done bool) {
	r.Lock()
	defer r.Unlock()
	if r.handling != nil {
		handling := *r.handling
		r.handling = nil
		handling <- done
	}
}

func (r *routeImpl) checkNotHandled() error {
	r.RLock()
	defer r.RUnlock()
	if r.handling == nil {
		return errors.New("Route is already handled!")
	}
	return nil
}

func (r *routeImpl) Request() Request {
	return fromChannel(r.initializer["request"]).(*requestImpl)
}

func unpackOptionalArgument(input interface{}) interface{} {
	inputValue := reflect.ValueOf(input)
	if inputValue.Kind() != reflect.Slice {
		panic("Needs to be a slice")
	}
	if inputValue.Len() == 0 {
		return nil
	}
	return inputValue.Index(0).Interface()
}

func (r *routeImpl) Abort(errorCode ...string) error {
	err := r.checkNotHandled()
	if err != nil {
		return err
	}
	err = r.raceWithPageClose(func() error {
		_, err := r.channel.Send("abort", map[string]interface{}{
			"errorCode":  unpackOptionalArgument(errorCode),
			"requestUrl": r.Request().(*requestImpl).initializer["url"],
		})
		return err
	})
	r.reportHandled(true)
	return err
}

func (r *routeImpl) raceWithPageClose(f func() error) error {
	page, ok := r.Request().Frame().Page().(*pageImpl)
	if !ok || page == nil {
		return f()
	}
	errChan := make(chan error, 1)
	go func() {
		errChan <- f()
	}()

	select {
	case <-page.closedOrCrashed:
		return errors.New("Page is closed or crashed")
	case err := <-errChan:
		return err
	}
}

func (r *routeImpl) Fulfill(options RouteFulfillOptions) error {
	err := r.checkNotHandled()
	if err != nil {
		return err
	}
	overrides := map[string]interface{}{
		"status": 200,
	}
	headers := make(map[string]string)

	if options.Response != nil {
		overrides["status"] = options.Response.Status()
		headers = options.Response.Headers()
		response, ok := options.Response.(*apiResponseImpl)
		if options.Body == nil && options.Path == nil && ok && response.request.connection == r.connection {
			overrides["fetchResponseUid"] = response.fetchUid()
		} else {
			options.Body, _ = options.Response.Body()
		}
		options.Response = nil
	}
	if options.Status != nil {
		overrides["status"] = *options.Status
		options.Status = nil
	}

	length := 0
	isBase64 := false
	var fileContentType string
	if _, ok := options.Body.(string); ok {
		isBase64 = false
	} else if body, ok := options.Body.([]byte); ok {
		options.Body = base64.StdEncoding.EncodeToString(body)
		length = len(body)
		isBase64 = true
	} else if options.Path != nil {
		content, err := os.ReadFile(*options.Path)
		if err != nil {
			return err
		}
		fileContentType = http.DetectContentType(content)
		options.Body = base64.StdEncoding.EncodeToString(content)
		isBase64 = true
		length = len(content)
	}

	if options.Headers != nil {
		headers = make(map[string]string)
		for key, val := range options.Headers {
			headers[strings.ToLower(key)] = val
		}
		options.Headers = nil
	}
	if options.ContentType != nil {
		headers["content-type"] = *options.ContentType
		options.ContentType = nil
	} else if options.Path != nil {
		headers["content-type"] = fileContentType
	}
	if _, ok := headers["content-length"]; !ok && length > 0 {
		headers["content-length"] = strconv.Itoa(length)
	}
	overrides["headers"] = serializeMapToNameAndValue(headers)
	overrides["isBase64"] = isBase64
	overrides["requestUrl"] = r.Request().(*requestImpl).initializer["url"]

	options.Path = nil
	err = r.raceWithPageClose(func() error {
		_, err := r.channel.Send("fulfill", options, overrides)
		return err
	})
	r.reportHandled(true)
	return err
}

func (r *routeImpl) Fallback(options ...RouteFallbackOptions) error {
	err := r.checkNotHandled()
	if err != nil {
		return err
	}
	opt := RouteFallbackOptions{}
	if len(options) == 1 {
		opt = options[0]
	}
	r.Request().(*requestImpl).applyFallbackOverrides(opt)
	r.reportHandled(false)
	return nil
}

func (r *routeImpl) Fetch(options ...RouteFetchOptions) (APIResponse, error) {
	request := r.Request().Frame().Page().Context().Request().(*apiRequestContextImpl)
	opt := APIRequestContextFetchOptions{}
	url := ""
	if len(options) > 0 {
		opt.Headers = options[0].Headers
		opt.Method = options[0].Method
		opt.Data = options[0].PostData
		if options[0].URL != nil {
			url = *options[0].URL
		}
	}
	return request.innerFetch(url, r.Request(), opt)
}

func (r *routeImpl) Continue(options ...RouteContinueOptions) error {
	option := &RouteFallbackOptions{}
	if len(options) == 1 {
		err := assignStructFields(option, options[0], true)
		if err != nil {
			return err
		}
	}

	err := r.checkNotHandled()
	if err != nil {
		return err
	}
	r.Request().(*requestImpl).applyFallbackOverrides(*option)
	err = r.internalContinue(false)
	r.reportHandled(true)
	return err
}

func (r *routeImpl) internalContinue(isInternal bool) error {
	overrides := make(map[string]interface{})
	overrides["url"] = r.Request().(*requestImpl).fallbackOverrides.URL
	overrides["method"] = r.Request().(*requestImpl).fallbackOverrides.Method
	headers := r.Request().(*requestImpl).fallbackOverrides.Headers
	if headers != nil {
		overrides["headers"] = serializeMapToNameAndValue(headers)
	}
	postDataBuf := r.Request().(*requestImpl).fallbackOverrides.PostDataBuffer
	if postDataBuf != nil {
		overrides["postData"] = base64.StdEncoding.EncodeToString(postDataBuf)
	}
	overrides["requestUrl"] = r.Request().(*requestImpl).initializer["url"]
	overrides["isFallback"] = isInternal
	_, err := r.channel.connection.WrapAPICall(func() (interface{}, error) {
		err := r.raceWithPageClose(func() error {
			_, err := r.channel.Send("continue", overrides)
			return err
		})
		return nil, err
	}, isInternal)
	return err
}

func (r *routeImpl) redirectedNavigationRequest(url string) error {
	err := r.checkNotHandled()
	if err != nil {
		return err
	}
	_, err = r.channel.Send("redirectNavigationRequest", map[string]interface{}{
		"url": url,
	})
	r.reportHandled(true)
	return err
}

func newRoute(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *routeImpl {
	bt := &routeImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
