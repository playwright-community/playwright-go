package playwright

type webSocketImpl struct {
	channelOwner
}

func (r *webSocketImpl) URL() string {
	return r.initializer["url"].(string)
}

func newWebsocket(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *webSocketImpl {
	resp := &webSocketImpl{}
	resp.createChannelOwner(resp, parent, objectType, guid, initializer)
	return resp
}
