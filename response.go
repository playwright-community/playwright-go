package playwright

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

func newResponse(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Response {
	resp := &Response{}
	resp.createChannelOwner(resp, parent, objectType, guid, initializer)
	return resp
}
