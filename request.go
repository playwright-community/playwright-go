package playwright

type Request struct {
	ChannelOwner
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

func newRequest(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Request {
	req := &Request{}
	req.createChannelOwner(req, parent, objectType, guid, initializer)
	return req
}
