package playwright

import (
	"encoding/base64"
	"encoding/json"
)

// RequestFailure represents a request failure
type RequestFailure struct {
	ErrorText string
}

// ResourceTiming represents the resource timing
type ResourceTiming struct {
	StartTime             float64
	DomainLookupStart     float64
	DomainLookupEnd       float64
	ConnectStart          float64
	SecureConnectionStart float64
	ConnectEnd            float64
	RequestStart          float64
	ResponseStart         float64
	ResponseEnd           float64
}

type requestImpl struct {
	channelOwner
	timing         *ResourceTiming
	headers        map[string]string
	redirectedFrom Request
	redirectedTo   Request
	failureText    string
}

func (r *requestImpl) URL() string {
	return r.initializer["url"].(string)
}

func (r *requestImpl) ResourceType() string {
	return r.initializer["resourceType"].(string)
}

func (r *requestImpl) Method() string {
	return r.initializer["method"].(string)
}

func (r *requestImpl) PostData() (string, error) {
	body, err := r.PostDataBuffer()
	if err != nil {
		return "", err
	}
	return string(body), err
}

func (r *requestImpl) PostDataJSON(v interface{}) error {
	body, err := r.PostDataBuffer()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func (r *requestImpl) PostDataBuffer() ([]byte, error) {
	if _, ok := r.initializer["postData"]; !ok {
		return nil, nil
	}
	return base64.StdEncoding.DecodeString(r.initializer["postData"].(string))
}

func (r *requestImpl) Headers() map[string]string {
	return r.headers
}

func (r *requestImpl) Response() (Response, error) {
	channel, err := r.channel.Send("response")
	if err != nil {
		return nil, err
	}
	channelOwner := fromNullableChannel(channel)
	if channelOwner == nil {
		return nil, nil
	}
	return channelOwner.(*responseImpl), nil
}

func (r *requestImpl) Frame() Frame {
	return fromChannel(r.initializer["frame"]).(*frameImpl)
}

func (r *requestImpl) IsNavigationRequest() bool {
	return r.initializer["isNavigationRequest"].(bool)
}

func (r *requestImpl) RedirectedFrom() Request {
	return r.redirectedFrom
}

func (r *requestImpl) RedirectedTo() Request {
	return r.redirectedTo
}

func (r *requestImpl) Failure() *RequestFailure {
	if r.failureText == "" {
		return nil
	}
	return &RequestFailure{
		ErrorText: r.failureText,
	}
}

func (r *requestImpl) Timing() *ResourceTiming {
	return r.timing
}

func newRequest(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *requestImpl {
	req := &requestImpl{}
	req.createChannelOwner(req, parent, objectType, guid, initializer)
	redirectedFrom := fromNullableChannel(initializer["redirectedFrom"])
	if redirectedFrom != nil {
		req.redirectedFrom = redirectedFrom.(*requestImpl)
	}
	if req.redirectedFrom != nil {
		req.redirectedFrom.(*requestImpl).redirectedTo = req
	}
	req.timing = &ResourceTiming{
		StartTime:             0,
		DomainLookupStart:     -1,
		DomainLookupEnd:       -1,
		ConnectStart:          -1,
		SecureConnectionStart: -1,
		ConnectEnd:            -1,
		RequestStart:          -1,
		ResponseStart:         -1,
		ResponseEnd:           -1,
	}
	req.headers = parseHeaders(req.initializer["headers"].([]interface{}))
	return req
}
