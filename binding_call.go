package playwright

type BindingCall struct {
	ChannelOwner
}

func newBindingCall(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *BindingCall {
	bt := &BindingCall{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
