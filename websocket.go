package playwright

type WebSocket struct {
	ChannelOwner
}

func (r *WebSocket) URL() string {
	return r.initializer["url"].(string)
}

func newWebsocket(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *WebSocket {
	resp := &WebSocket{}
	resp.createChannelOwner(resp, parent, objectType, guid, initializer)
	return resp
}
