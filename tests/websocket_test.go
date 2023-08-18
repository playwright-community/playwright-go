package playwright_test

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"

	goContext "context"

	"github.com/gorilla/websocket"
	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

type webSocketServer struct {
	PORT   int
	server *http.Server
}

func (ws *webSocketServer) Stop() {
	if err := ws.server.Shutdown(goContext.Background()); err != nil {
		log.Printf("could not stop ws server: %v", err)
	}
}

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (ws *webSocketServer) handler(w http.ResponseWriter, r *http.Request) {
	c, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("could not upgrade ws connection:", err)
		return
	}
	if err := c.WriteMessage(websocket.TextMessage, []byte("incoming")); err != nil {
		log.Println("could not write ws message:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNoStatusReceived) {
				log.Println("could not read ws message:", err)
			}
			break
		}
		if bytes.Equal(message, []byte("echo-bin")) {
			if err := c.WriteMessage(websocket.BinaryMessage, []byte{4, 2}); err != nil {
				log.Println("could not write ws message:", err)
				return
			}
		}
		if bytes.Equal(message, []byte("echo-text")) {
			if err := c.WriteMessage(mt, []byte("text")); err != nil {
				log.Println("could not write ws message:", err)
				return
			}
		}
	}
}

func newWebsocketServer() *webSocketServer {
	wsServer := &webSocketServer{
		PORT: 8012,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsServer.handler)
	wsServer.server = &http.Server{Handler: mux}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", wsServer.PORT))
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		if err := wsServer.server.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Printf("could not start ws server: %v", err)
		}
	}()
	return wsServer
}

func TestWebSocketShouldWork(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	wsServer := newWebsocketServer()
	defer wsServer.Stop()
	value, err := page.Evaluate(`port => {
        let cb;
        const result = new Promise(f => cb = f);
        const ws = new WebSocket('ws://localhost:' + port + '/ws');
        ws.addEventListener('message', data => { ws.close(); cb(data.data); });
        return result;
	}`, wsServer.PORT)
	require.NoError(t, err)
	require.Equal(t, "incoming", value)
}

func TestWebSocketShouldEmitCloseEvents(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	wsServer := newWebsocketServer()
	defer wsServer.Stop()
	ws, err := page.ExpectWebSocket(func() error {
		_, err := page.Evaluate(`port => {
					const ws = new WebSocket('ws://localhost:' + port + '/ws');
					ws.addEventListener('message', data => { ws.close() });
        }`, wsServer.PORT)
		return err
	}, playwright.PageExpectWebSocketOptions{
		Timeout: playwright.Float(1000),
	})
	require.NoError(t, err)
	ws.OnClose(func(w playwright.WebSocket) {
		require.Equal(t, w.URL(), fmt.Sprintf("ws://localhost:%d/ws", wsServer.PORT))
	})
	require.Equal(t, ws.URL(), fmt.Sprintf("ws://localhost:%d/ws", wsServer.PORT))
	if !ws.IsClosed() {
		_, err = ws.WaitForEvent("close")
		require.NoError(t, err)
	}
	require.True(t, ws.IsClosed())
}

func TestWebSocketShouldEmitFrameEvents(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	wsServer := newWebsocketServer()
	defer wsServer.Stop()

	sent := [][]byte{}
	received := [][]byte{}

	page.OnWebSocket(func(ws playwright.WebSocket) {
		ws.OnFrameSent(func(payload []byte) {
			sent = append(sent, payload)
		})
		ws.OnFrameReceived(func(payload []byte) {
			received = append(received, payload)
		})
	})
	ws, err := page.ExpectWebSocket(func() error {
		_, err := page.Evaluate(`port => {
						let count = 0;
            const ws = new WebSocket('ws://localhost:' + port + '/ws');
            ws.addEventListener('open', () => {
                ws.send('echo-text');
            });
						ws.addEventListener('message', data => { count++; if (count >= 2) { ws.close() } });
        }`, wsServer.PORT)
		return err
	})
	require.NoError(t, err)
	if !ws.IsClosed() {
		_, err = ws.WaitForEvent("close", playwright.WebSocketWaitForEventOptions{
			Timeout: playwright.Float(1000),
		})
		require.NoError(t, err)
	}

	require.Equal(t, sent, [][]byte{[]byte("echo-text")})
	require.Equal(t, received, [][]byte{[]byte("incoming"), []byte("text")})
}

func TestWebSocketShouldEmitBinaryFrameEvents(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	wsServer := newWebsocketServer()
	defer wsServer.Stop()

	sent := [][]byte{}
	received := [][]byte{}

	page.Once("websocket", func(ws playwright.WebSocket) {
		ws.OnFrameSent(func(payload []byte) {
			sent = append(sent, payload)
		})
		ws.OnFrameReceived(func(payload []byte) {
			received = append(received, payload)
		})
	})
	ws, err := page.ExpectWebSocket(func() error {
		_, err := page.Evaluate(`port => {
						let count = 0;
            const ws = new WebSocket('ws://localhost:' + port + '/ws');
            ws.addEventListener('open', () => {
                const binary = new Uint8Array(5);
                for (let i = 0; i < 5; ++i)
                    binary[i] = i;
                ws.send(binary);
                ws.send('echo-bin');
            });
						ws.addEventListener('message', data => { count++; if (count >= 2) { ws.close() } });
        }`, wsServer.PORT)
		return err
	})
	require.NoError(t, err)
	if !ws.IsClosed() {
		_, err = ws.WaitForEvent("close")
		require.NoError(t, err)
	}

	require.Equal(t, sent, [][]byte{{0, 1, 2, 3, 4}, []byte("echo-bin")})
	require.Equal(t, received, [][]byte{[]byte("incoming"), {4, 2}})
}

func TestWebSocketShouldRejectWaitForEventOnCloseAndError(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	wsServer := newWebsocketServer()
	defer wsServer.Stop()
	ws, err := page.ExpectWebSocket(func() error {
		_, err := page.Evaluate(`port => {
            ws = new WebSocket('ws://localhost:' + port + '/ws');
        }`, wsServer.PORT)
		return err
	})
	require.NoError(t, err)
	// event may have been generated before interception
	_, err = ws.WaitForEvent("framereceived")
	require.NoError(t, err)
	_, err = ws.ExpectEvent("framesent", func() error {
		_, err := page.Evaluate(`window.ws.close()`)
		return err
	}, playwright.WebSocketExpectEventOptions{
		Timeout: playwright.Float(1000),
		Predicate: func(ev interface{}) bool {
			return true
		},
	})
	require.ErrorContains(t, err, "websocket closed")
}

func TestWebSocketShouldEmitErrorEvent(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	wsServer := newWebsocketServer()
	defer wsServer.Stop()
	chanMsg := make(chan string, 1)
	page.OnWebSocket(func(ws playwright.WebSocket) {
		ws.OnSocketError(func(err string) {
			chanMsg <- err
		})
	})
	_, err := page.Evaluate(`port => {
			ws = new WebSocket('ws://localhost:' + port + '/bogus-ws');
		}`, wsServer.PORT)
	require.NoError(t, err)
	msg := <-chanMsg
	if isFirefox {
		require.Equal(t, msg, "CLOSE_ABNORMAL")
	} else {
		require.Contains(t, msg, ": 404")
	}
}
