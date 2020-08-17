package playwright

type Response struct {
	ChannelOwner
}

func newResponse(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Response {
	bt := &Response{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
