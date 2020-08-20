package playwright

type BrowserServer struct {
	ChannelOwner
}

func newBrowserServer(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *BrowserServer {
	bt := &BrowserServer{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
