package playwright

import (
	"encoding/base64"
	"encoding/json"
)

type responseImpl struct {
	channelOwner
	request *requestImpl
}

func (r *responseImpl) URL() string {
	return r.initializer["url"].(string)
}

func (r *responseImpl) Ok() bool {
	return r.Status() == 0 || (r.Status() >= 200 && r.Status() <= 299)
}

func (r *responseImpl) Status() int {
	return int(r.initializer["status"].(float64))
}

func (r *responseImpl) StatusText() string {
	return r.initializer["statusText"].(string)
}

func (r *responseImpl) Headers() map[string]string {
	return parseHeaders(r.initializer["headers"].([]interface{}))
}

func (r *responseImpl) Finished() error {
	_, err := r.channel.Send("finished")
	return err
}

func (r *responseImpl) Body() ([]byte, error) {
	b64Body, err := r.channel.Send("body")
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(b64Body.(string))
}

func (r *responseImpl) Text() (string, error) {
	body, err := r.Body()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (r *responseImpl) JSON(v interface{}) error {
	body, err := r.Body()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func (r *responseImpl) Request() Request {
	return r.request
}

func (r *responseImpl) Frame() Frame {
	return r.request.Frame()
}

func newResponse(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *responseImpl {
	resp := &responseImpl{}
	resp.createChannelOwner(resp, parent, objectType, guid, initializer)
	timing := resp.initializer["timing"].(map[string]interface{})
	resp.request = fromChannel(resp.initializer["request"]).(*requestImpl)
	resp.request.timing = &ResourceTiming{
		StartTime:             timing["startTime"].(float64),
		DomainLookupStart:     timing["domainLookupStart"].(float64),
		DomainLookupEnd:       timing["domainLookupEnd"].(float64),
		ConnectStart:          timing["connectStart"].(float64),
		SecureConnectionStart: timing["secureConnectionStart"].(float64),
		ConnectEnd:            timing["connectEnd"].(float64),
		RequestStart:          timing["requestStart"].(float64),
		ResponseStart:         timing["responseStart"].(float64),
	}
	resp.request.headers = parseHeaders(resp.initializer["requestHeaders"].([]interface{}))
	return resp
}
