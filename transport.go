package playwright

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/go-jose/go-jose/v3/json"
)

type pipeTransport struct {
	stdin     io.WriteCloser
	stdout    io.ReadCloser
	onmessage func(msg *message)
	rLock     sync.Mutex
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
			fmt.Fprint(os.Stdout, "\x1b[33mRECV>\x1b[0m\n")
			if err := json.NewEncoder(os.Stdout).Encode(msg); err != nil {
				log.Printf("could not encode json: %v", err)
			}
		}
		t.onmessage(msg)
	}
}

type errorPayload struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Stack   string `json:"stack"`
}

type message struct {
	ID     int                    `json:"id"`
	GUID   string                 `json:"guid"`
	Method string                 `json:"method,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
	Result interface{}            `json:"result,omitempty"`
	Error  *struct {
		Error errorPayload `json:"error"`
	} `json:"error,omitempty"`
}

func (t *pipeTransport) Send(message map[string]interface{}) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("could not marshal json: %w", err)
	}
	if os.Getenv("DEBUGP") != "" {
		fmt.Fprint(os.Stdout, "\x1b[32mSEND>\x1b[0m\n")
		if err := json.NewEncoder(os.Stdout).Encode(message); err != nil {
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

func newPipeTransport(stdin io.WriteCloser, stdout io.ReadCloser) *pipeTransport {
	return &pipeTransport{
		stdout: stdout,
		stdin:  stdin,
	}
}
