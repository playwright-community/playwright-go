package playwright

type Request struct {
	ChannelOwner
}

func newRequest(parent *ChannelOwner, objectType string, guid string, initializer interface{}) *Request {
	bt := &Request{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
