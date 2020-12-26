package playwright

type webSocketImpl struct {
	channelOwner
}

func (w *webSocketImpl) URL() string {
	return w.initializer["url"].(string)
}

func newWebsocket(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *webSocketImpl {
	resp := &webSocketImpl{}
	resp.createChannelOwner(resp, parent, objectType, guid, initializer)
	return resp
}
