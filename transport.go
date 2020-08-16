package playwright

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/square/go-jose.v2/json"
)

type Transport struct {
	stdin    io.WriteCloser
	stdout   io.ReadCloser
	dispatch func(msg *Message) error
}

func (t *Transport) SetDispatch(dispatch func(msg *Message) error) {
	t.dispatch = dispatch
}

func (t *Transport) Start() error {
	reader := bufio.NewReader(t.stdout)
	for {
		lengthContent := make([]byte, 4)
		_, err := io.ReadFull(reader, lengthContent)
		if err != nil {
			return fmt.Errorf("could not read padding: %v", err)
		}
		length := binary.LittleEndian.Uint32(lengthContent)

		msg := &Message{}
		if err := json.NewDecoder(io.LimitReader(reader, int64(length))).Decode(&msg); err != nil {
			return fmt.Errorf("could not parse json: %v", err)
		}
		if os.Getenv("DEBUGP") != "" {
			spew.Dump("RECV>", msg)
		}
		t.dispatch(msg)
	}
}

func (t *Transport) Stop() error {
	return nil
}

type Message struct {
	ID     int    `json:"id"`
	GUID   string `json:"guid"`
	Method string `json:"method"`
	Params struct {
		Type        string      `json:"type"`
		GUID        string      `json:"guid"`
		Initializer interface{} `json:"initializer"`
	} `json:"params"`
	Result interface{} `json:"result"`
	Error  struct {
		Message string `json:"message"`
		Stack   string `json:"stack"`
	} `json:"error"`
}

func (t *Transport) Send(message map[string]interface{}) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("could not marshal json: %v", err)
	}
	if os.Getenv("DEBUGP") != "" {
		spew.Dump("SEND>", message)
	}
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(len(msg)))
	t.stdin.Write(b)
	t.stdin.Write(msg)
	return nil
}

func newTransport(stdin io.WriteCloser, stdout io.ReadCloser) *Transport {
	return &Transport{
		stdout: stdout,
		stdin:  stdin,
	}
}
