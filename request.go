package playwright

type Request struct {
	ChannelOwner
}

func newRequest(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Request {
	bt := &Request{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
