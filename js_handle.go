package playwright

type JSHandle struct {
	ChannelOwner
}

func newJSHandle(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *JSHandle {
	bt := &JSHandle{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
