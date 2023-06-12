package playwright

import (
	"encoding/base64"
	"errors"
	"log"
)

type webSocketImpl struct {
	channelOwner
	isClosed bool
	page     *pageImpl
}

func (ws *webSocketImpl) URL() string {
	return ws.initializer["url"].(string)
}

func newWebsocket(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *webSocketImpl {
	ws := &webSocketImpl{}
	ws.createChannelOwner(ws, parent, objectType, guid, initializer)
	ws.page = fromChannel(parent.channel).(*pageImpl)
	ws.channel.On("close", func() {
		ws.Lock()
		ws.isClosed = true
		ws.Unlock()
		ws.Emit("close")
	})
	ws.channel.On(
		"frameSent",
		func(params map[string]interface{}) {
			ws.onFrameSent(params["opcode"].(float64), params["data"].(string))
		},
	)
	ws.channel.On(
		"frameReceived",
		func(params map[string]interface{}) {
			ws.onFrameReceived(params["opcode"].(float64), params["data"].(string))
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

func (ws *webSocketImpl) onFrameSent(opcode float64, data string) {
	if opcode == 2 {
		payload, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			log.Printf("could not decode WebSocket.onFrameSent payload: %v", err)
			return
		}
		ws.Emit("framesent", payload)
	} else {
		ws.Emit("framesent", []byte(data))
	}
}

func (ws *webSocketImpl) onFrameReceived(opcode float64, data string) {
	if opcode == 2 {
		payload, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			log.Printf("could not decode WebSocket.onFrameReceived payload: %v", err)
			return
		}
		ws.Emit("framereceived", payload)
	} else {
		ws.Emit("framereceived", []byte(data))
	}
}

func (ws *webSocketImpl) ExpectEvent(event string, cb func() error, options ...WebSocketWaitForEventOptions) (interface{}, error) {
	return ws.expectEvent(event, cb, options...)
}

func (ws *webSocketImpl) WaitForEvent(event string, options ...WebSocketWaitForEventOptions) (interface{}, error) {
	return ws.expectEvent(event, nil, options...)
}

func (ws *webSocketImpl) expectEvent(event string, cb func() error, options ...WebSocketWaitForEventOptions) (interface{}, error) {
	var predicate interface{} = nil
	var timeout = ws.page.timeoutSettings.Timeout()
	if len(options) == 1 {
		if options[0].Timeout != nil {
			timeout = *options[0].Timeout
		}
		if options[0].Predicate != nil {
			predicate = options[0].Predicate
		}
	}
	waiter := newWaiter().WithTimeout(timeout)
	if event != "close" {
		waiter.RejectOnEvent(ws, "close", errors.New("websocket closed"))
	}
	if event != "error" {
		waiter.RejectOnEvent(ws, "error", errors.New("websocket error"))
	}
	if cb == nil {
		return waiter.WaitForEvent(ws, event, predicate).Wait()
	} else {
		return waiter.WaitForEvent(ws, event, predicate).RunAndWait(cb)
	}
}

func (ws *webSocketImpl) IsClosed() bool {
	ws.RLock()
	defer ws.RUnlock()
	return ws.isClosed
}
