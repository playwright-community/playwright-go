package playwright

type webSocketImpl struct {
	channelOwner
	isClosed bool
}

func (w *webSocketImpl) URL() string {
	return w.initializer["url"].(string)
}

func newWebsocket(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *webSocketImpl {
	resp := &webSocketImpl{}
	resp.createChannelOwner(resp, parent, objectType, guid, initializer)
	resp.channel.On("close", func(ev map[string]interface{}) {
		resp.isClosed = true
		resp.Emit("close")
	})
	return resp
}

func (w *webSocketImpl) IsClosed() bool {
	return w.isClosed
}
