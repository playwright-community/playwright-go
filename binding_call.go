package playwright

type BindingCall struct {
	channelOwner
}

func newBindingCall(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *BindingCall {
	bt := &BindingCall{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
