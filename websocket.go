package playwright

import (
	"encoding/base64"
	"log"
	"reflect"
)

type webSocketImpl struct {
	channelOwner
	isClosed bool
}

func (ws *webSocketImpl) URL() string {
	return ws.initializer["url"].(string)
}

func newWebsocket(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *webSocketImpl {
	ws := &webSocketImpl{}
	ws.createChannelOwner(ws, parent, objectType, guid, initializer)
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

func (ws *webSocketImpl) WaitForEvent(event string, predicate ...interface{}) interface{} {
	return <-ws.WaitForEventCh(event, predicate...)
}

func (ws *webSocketImpl) WaitForEventCh(event string, predicate ...interface{}) <-chan interface{} {
	evChan := make(chan interface{}, 1)
	ws.Once(event, func(ev ...interface{}) {
		if len(predicate) == 0 {
			if len(ev) == 1 {
				evChan <- ev[0]
			} else {
				evChan <- nil
			}
		} else if len(predicate) == 1 {
			result := reflect.ValueOf(predicate[0]).Call([]reflect.Value{reflect.ValueOf(ev[0])})
			if result[0].Bool() {
				evChan <- ev[0]
			}
		}
	})
	return evChan
}

func (w *webSocketImpl) IsClosed() bool {
	w.RLock()
	defer w.RUnlock()
	return w.isClosed
}
