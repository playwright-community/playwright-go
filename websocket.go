package playwright

import (
	"encoding/base64"
	"log"
)

type webSocketImpl struct {
	channelOwner
	isClosed bool
}

func (w *webSocketImpl) URL() string {
	return w.initializer["url"].(string)
}

func newWebsocket(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *webSocketImpl {
	ws := &webSocketImpl{}
	ws.createChannelOwner(ws, parent, objectType, guid, initializer)
	ws.channel.On("close", func() {
		ws.isClosed = true
		ws.Emit("close")
	})
	ws.channel.On(
		"frameSent",
		func(params map[string]interface{}) {
			ws.onFrameSent(params["opcode"].(int), params["data"].(string))
		},
	)
	ws.channel.On(
		"frameReceived",
		func(params map[string]interface{}) {
			ws.onFrameReceived(params["opcode"].(int), params["data"].(string))
		},
	)
	ws.channel.On(
		"error",
		func(params map[string]interface{}) {
			ws.Emit("error", params["error"])
		},
	)
	return ws
}

func (ws *webSocketImpl) onFrameSent(opcode int, data string) {
	if opcode == 2 {
		payload, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			log.Printf("could not decode WebSocket.onFrameSent payload: %v", err)
			return
		}
		ws.Emit("framesent", payload)
	} else {
		ws.Emit("framesent", data)
	}
}

func (ws *webSocketImpl) onFrameReceived(opcode int, data string) {
	if opcode == 2 {
		payload, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			log.Printf("could not decode WebSocket.onFrameReceived payload: %v", err)
			return
		}
		ws.Emit("framereceived", payload)
	} else {
		ws.Emit("framereceived", data)
	}
}

func (w *webSocketImpl) IsClosed() bool {
	return w.isClosed
}
