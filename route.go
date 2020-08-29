package playwright

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Route struct {
	ChannelOwner
}

func (r *Route) Request() *Request {
	return fromChannel(r.initializer["request"]).(*Request)
}

func (r *Route) Abort(errorCode *string) error {
	_, err := r.channel.Send("abort", map[string]*string{
		"errorCode": errorCode,
	})
	return err
}

type RouteFulfillOptions struct {
	Status      *int              `json:"status"`
	Headers     map[string]string `json:"headers"`
	Body        interface{}       `json:"body"`
	Path        *string           `json:"path"`
	ContentType *string           `json:"contentType"`
}

func (r *Route) Fulfill(options RouteFulfillOptions) error {
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
		content, err := ioutil.ReadFile(*options.Path)
		if err != nil {
			return err
		}
		fileContentType = http.DetectContentType(content)
		options.Body = base64.StdEncoding.EncodeToString(content)
		isBase64 = true
		length = len(content)
		options.Path = nil
	}

	headers := make(map[string]string)
	if options.Headers != nil {
		for key, val := range options.Headers {
			headers[strings.ToLower(key)] = val
		}
		if options.ContentType != nil {
			headers["content-type"] = *options.ContentType
		} else if options.Path != nil {
			headers["content-type"] = fileContentType
		}
		if _, ok := headers["content-length"]; !ok && length > 0 {
			headers["content-length"] = strconv.Itoa(length)
		}
	}

	_, err := r.channel.Send("fulfill", map[string]interface{}{
		"isBase64": isBase64,
		"headers":  serializeHeaders(headers),
	}, options)
	return err
}

type RouteContinueOptions struct {
	Method   *string           `json:"method"`
	Headers  map[string]string `json:"headers"`
	PostData interface{}       `json:"postData"`
}

func (r *Route) Continue(options RouteContinueOptions) error {
	overrides := make(map[string]interface{})
	if options.Method != nil {
		overrides["method"] = options.Method
	}
	if options.Headers != nil {
		overrides["headers"] = serializeHeaders(options.Headers)
	}
	if options.PostData != nil {
		switch v := options.PostData.(type) {
		case string:
			overrides["postData"] = base64.StdEncoding.EncodeToString([]byte(v))
		case []byte:
			overrides["postData"] = base64.StdEncoding.EncodeToString(v)
		}
	}
	_, err := r.channel.Send("continue", overrides)
	return err
}

func newRoute(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Route {
	bt := &Route{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
