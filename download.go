package playwright

type Download struct {
	ChannelOwner
}

func (d *Download) String() string {
	return d.SuggestedFilename()
}

func (d *Download) URL() string {
	return d.initializer["url"].(string)
}

func (d *Download) SuggestedFilename() string {
	return d.initializer["suggestedFilename"].(string)
}

func newDownload(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Download {
	bt := &Download{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
