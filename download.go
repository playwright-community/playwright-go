package playwright

type Download struct {
	ChannelOwner
}

func newDownload(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Download {
	bt := &Download{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
