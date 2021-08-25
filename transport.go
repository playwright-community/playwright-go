package playwright

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"gopkg.in/square/go-jose.v2/json"
)

type transport interface {
	Start() error
	Stop() error
	Send(message map[string]interface{}) error
	SetDispatch(dispatch func(msg *message))
}

type pipeTransport struct {
	stdin    io.WriteCloser
	stdout   io.ReadCloser
	dispatch func(msg *message)
	rLock    sync.Mutex
}

type webSocketTransport struct {
	OnClose  func()
	url      string
	conn     *websocket.Conn
	options  BrowserTypeConnectOptions
	dispatch func(msg *message)
	stopped  atomic.Value
}

func (t *webSocketTransport) Start() error {
	var headers http.Header
	if t.options.Headers != nil {
		for k, v := range t.options.Headers {
			headers.Add(k, v)
		}
	}
	conn, _, err := websocket.DefaultDialer.Dial(t.url, headers)
	if err != nil {
		return fmt.Errorf("could not connect to websocket: %v", err)
	}
	t.conn = conn
	for !t.stopped.Load().(bool) {
		if t.options.Timeout != nil {
			if err := t.conn.SetReadDeadline(time.Now().Add(time.Duration(*t.options.Timeout))); err != nil {
				return err
			}
		}
		if t.options.SlowMo != nil {
			time.Sleep(time.Duration(*t.options.SlowMo) * time.Millisecond)
		}
		msg := &message{}
		err := t.conn.ReadJSON(msg)
		if err != nil {
			_ = t.Stop()
			return nil
		}
		t.dispatch(msg)
	}
	return nil
}

func (t *webSocketTransport) Send(message map[string]interface{}) error {
	if err := t.conn.WriteJSON(message); err != nil {
		t.stopped.Store(true)
		return fmt.Errorf("could not write json: %w", err)
	}
	return nil
}

func (t *webSocketTransport) SetDispatch(dispatch func(msg *message)) {
	t.dispatch = dispatch
}

func (t *webSocketTransport) Stop() error {
	t.stopped.Store(true)
	if err := t.conn.Close(); err != nil {
		return fmt.Errorf("could not close websocket: %w", err)
	}
	if t.OnClose != nil {
		t.OnClose()
	}
	return nil
}

func (t *pipeTransport) Start() error {
	reader := bufio.NewReader(t.stdout)
	for {
		lengthContent := make([]byte, 4)
		_, err := io.ReadFull(reader, lengthContent)
		if err == io.EOF || errors.Is(err, os.ErrClosed) {
			return nil
		} else if err != nil {
			return fmt.Errorf("could not read padding: %w", err)
		}
		length := binary.LittleEndian.Uint32(lengthContent)

		msg := &message{}
		if err := json.NewDecoder(io.LimitReader(reader, int64(length))).Decode(&msg); err != nil {
			return fmt.Errorf("could not decode json: %w", err)
		}
		if os.Getenv("DEBUGP") != "" {
			fmt.Print("RECV>")
			if err := json.NewEncoder(os.Stderr).Encode(msg); err != nil {
				log.Printf("could not encode json: %v", err)
			}
		}
		t.dispatch(msg)
	}
}

func (t *pipeTransport) SetDispatch(dispatch func(msg *message)) {
	t.dispatch = dispatch
}
func (t *pipeTransport) Stop() error {
	return nil
}

type errorPayload struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Stack   string `json:"stack"`
}

type message struct {
	ID     int                    `json:"id"`
	GUID   string                 `json:"guid"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
	Result interface{}            `json:"result"`
	Error  *struct {
		Error errorPayload `json:"error"`
	} `json:"error"`
}

func (t *pipeTransport) Send(message map[string]interface{}) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("could not marshal json: %w", err)
	}
	if os.Getenv("DEBUGP") != "" {
		fmt.Print("SEND>")
		if err := json.NewEncoder(os.Stderr).Encode(message); err != nil {
			log.Printf("could not encode json: %v", err)
		}
	}
	lengthPadding := make([]byte, 4)
	t.rLock.Lock()
	defer t.rLock.Unlock()
	binary.LittleEndian.PutUint32(lengthPadding, uint32(len(msg)))
	if _, err = t.stdin.Write(lengthPadding); err != nil {
		return err
	}
	if _, err = t.stdin.Write(msg); err != nil {
		return err
	}

	return nil
}

func newPipeTransport(stdin io.WriteCloser, stdout io.ReadCloser) transport {
	return &pipeTransport{
		stdout: stdout,
		stdin:  stdin,
	}
}
func newWebSocketTransport(url string, options ...BrowserTypeConnectOptions) transport {
	var option BrowserTypeConnectOptions
	if len(options) == 1 {
		option = options[0]
	} else {
		option = BrowserTypeConnectOptions{}
	}
	t := &webSocketTransport{
		url:     url,
		options: option,
	}
	t.stopped.Store(false)
	return t
}
