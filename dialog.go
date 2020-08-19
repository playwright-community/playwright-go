package playwright

type Dialog struct {
	ChannelOwner
}

func newDialog(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Dialog {
	bt := &Dialog{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
