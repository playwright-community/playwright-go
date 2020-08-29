package playwright

import (
	"encoding/base64"
	"encoding/json"
)

type RequestFailure struct {
	ErrorText string
}

type Request struct {
	ChannelOwner
	redirectedFrom *Request
	redirectedTo   *Request
	failureText    string
}

func (r *Request) URL() string {
	return r.initializer["url"].(string)
}

func (r *Request) ResourceType() string {
	return r.initializer["resourceType"].(string)
}

func (r *Request) Method() string {
	return r.initializer["method"].(string)
}

func (r *Request) PostDataRaw() ([]byte, error) {
	if _, ok := r.initializer["postData"]; !ok {
		return []byte{}, nil
	}
	return base64.StdEncoding.DecodeString(r.initializer["postData"].(string))
}

func (r *Request) PostData() (string, error) {
	body, err := r.PostDataRaw()
	if err != nil {
		return "", err
	}
	return string(body), err
}

func (r *Request) PostDataJSON(v interface{}) error {
	body, err := r.PostDataRaw()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func (r *Request) Headers() map[string]string {
	return parseHeaders(r.initializer["headers"].([]interface{}))
}

func (r *Request) Response() (*Response, error) {
	resp, err := r.channel.Send("response")
	if err != nil {
		return nil, err
	}
	resp = fromNullableChannel(resp)
	if resp == nil {
		return nil, nil
	}
	return resp.(*Response), nil
}

func (r *Request) Frame() *Frame {
	return fromChannel(r.initializer["frame"]).(*Frame)
}

func (r *Request) IsNavigationRequest() bool {
	return r.initializer["isNavigationRequest"].(bool)
}

func (r *Request) RedirectedFrom() *Request {
	return r.redirectedFrom
}

func (r *Request) RedirectedTo() *Request {
	return r.redirectedTo
}

func (r *Request) Failure() *RequestFailure {
	if r.failureText == "" {
		return nil
	}
	return &RequestFailure{
		ErrorText: r.failureText,
	}
}

func newRequest(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Request {
	req := &Request{}
	req.createChannelOwner(req, parent, objectType, guid, initializer)
	redirectedFrom := fromNullableChannel(initializer["redirectedFrom"])
	if redirectedFrom != nil {
		req.redirectedFrom = redirectedFrom.(*Request)
	}
	if req.redirectedFrom != nil {
		req.redirectedFrom.redirectedTo = req
	}
	return req
}
