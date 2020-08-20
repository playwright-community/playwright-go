package playwright

type Route struct {
	ChannelOwner
}

func newRoute(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Route {
	bt := &Route{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
