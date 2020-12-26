package playwright

import (
	"encoding/base64"
	"encoding/json"
)

type Response struct {
	ChannelOwner
}

func (r *Response) URL() string {
	return r.initializer["url"].(string)
}

func (r *Response) Ok() bool {
	return r.Status() == 0 || (r.Status() >= 200 && r.Status() <= 299)
}

func (r *Response) Status() int {
	return int(r.initializer["status"].(float64))
}

func (r *Response) StatusText() string {
	return r.initializer["statusText"].(string)
}

func (r *Response) Headers() map[string]string {
	return parseHeaders(r.initializer["headers"].([]interface{}))
}

func (r *Response) Finished() error {
	_, err := r.channel.Send("finished")
	return err
}

func (r *Response) Body() ([]byte, error) {
	b64Body, err := r.channel.Send("body")
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(b64Body.(string))
}

func (r *Response) Text() (string, error) {
	body, err := r.Body()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (r *Response) JSON(v interface{}) error {
	body, err := r.Body()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func (r *Response) Request() RequestI {
	return fromChannel(r.initializer["request"]).(*Request)
}

func (r *Response) Frame() FrameI {
	return r.Request().Frame()
}

func newResponse(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Response {
	resp := &Response{}
	resp.createChannelOwner(resp, parent, objectType, guid, initializer)
	return resp
}
