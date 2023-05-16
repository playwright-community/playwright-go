package playwright

type localUtilsImpl struct {
	channelOwner
}

func newLocalUtils(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *localUtilsImpl {
	l := &localUtilsImpl{}
	l.createChannelOwner(l, parent, objectType, guid, initializer)
	return l
}
