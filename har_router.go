package playwright

import (
	"errors"
	"log"
)

type harRouter struct {
	localUtils     *localUtilsImpl
	harId          string
	notFoundAction HarNotFound
	urlOrPredicate interface{}
	err            error
}

func (r *harRouter) addContextRoute(context BrowserContext) error {
	if r.err != nil {
		return r.err
	}
	err := context.Route(r.urlOrPredicate, func(route Route) {
		err := r.handle(route)
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		return err
	}
	return r.err
}

func (r *harRouter) addPageRoute(page Page) error {
	if r.err != nil {
		return r.err
	}
	err := page.Route(r.urlOrPredicate, func(route Route) {
		err := r.handle(route)
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		return err
	}
	return r.err
}

func (r *harRouter) dispose() {
	go func() {
		r.err = r.localUtils.HarClose(r.harId)
	}()
}

func (r *harRouter) handle(route Route) error {
	if r.err != nil {
		return r.err
	}
	request := route.Request()
	postData, err := request.PostDataBuffer()
	if err != nil {
		return err
	}
	response, err := r.localUtils.HarLookup(harLookupOptions{
		HarId:               r.harId,
		URL:                 request.URL(),
		Method:              request.Method(),
		Headers:             request.Headers(),
		IsNavigationRequest: request.IsNavigationRequest(),
		PostData:            postData,
	})
	if err != nil {
		return err
	}
	switch response.Action {
	case "redirect":
		if response.RedirectURL == nil {
			return errors.New("redirect url is null")
		}
		return route.(*routeImpl).redirectedNavigationRequest(*response.RedirectURL)
	case "fulfill":
		if response.Body == nil {
			return errors.New("fulfill body is null")
		}
		return route.Fulfill(RouteFulfillOptions{
			Body:    *response.Body,
			Status:  response.Status,
			Headers: deserializeNameAndValueToMap(response.Headers),
		})
	case "error":
		log.Printf("har action error: %v", *response.Message)
		fallthrough
	case "noentry":
	}
	if r.notFoundAction == *HarNotFoundAbort {
		return route.Abort()
	}
	return route.Fallback()
}

func newHarRouter(localUtils *localUtilsImpl, file string, notFoundAction HarNotFound, urlOrPredicate interface{}) *harRouter {
	harId, err := localUtils.HarOpen(file)
	var url interface{} = "**/*"
	if urlOrPredicate != nil {
		url = urlOrPredicate
	}
	return &harRouter{
		localUtils:     localUtils,
		harId:          harId,
		notFoundAction: notFoundAction,
		urlOrPredicate: url,
		err:            err,
	}
}
