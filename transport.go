package playwright

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/go-jose/go-jose/v3/json"
)

type transport interface {
	Send(msg map[string]interface{}) error
	Poll() (*message, error)
	Close() error
}

type pipeTransport struct {
	writer    io.WriteCloser
	bufReader *bufio.Reader
	closed    chan struct{}
	onClose   func() error
}

func (t *pipeTransport) Poll() (*message, error) {
	if t.isClosed() {
		return nil, fmt.Errorf("transport closed")
	}
	lengthContent := make([]byte, 4)
	_, err := io.ReadFull(t.bufReader, lengthContent)
	if err == io.EOF || errors.Is(err, os.ErrClosed) {
		return nil, fmt.Errorf("pipe closed: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("could not read padding: %w", err)
	}
	length := binary.LittleEndian.Uint32(lengthContent)

	msg := &message{}
	if err := json.NewDecoder(io.LimitReader(t.bufReader, int64(length))).Decode(&msg); err != nil {
		return nil, fmt.Errorf("could not decode json: %w", err)
	}
	if os.Getenv("DEBUGP") != "" {
		fmt.Fprint(os.Stdout, "\x1b[33mRECV>\x1b[0m\n")
		if err := json.NewEncoder(os.Stdout).Encode(msg); err != nil {
			log.Printf("could not encode json: %v", err)
		}
	}
	return msg, nil
}

type message struct {
	ID     int                    `json:"id"`
	GUID   string                 `json:"guid"`
	Method string                 `json:"method,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
	Result interface{}            `json:"result,omitempty"`
	Error  *struct {
		Error Error `json:"error"`
	} `json:"error,omitempty"`
}

func (t *pipeTransport) Send(msg map[string]interface{}) error {
	if t.isClosed() {
		return fmt.Errorf("transport closed")
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("pipeTransport: could not marshal json: %w", err)
	}
	if os.Getenv("DEBUGP") != "" {
		fmt.Fprint(os.Stdout, "\x1b[32mSEND>\x1b[0m\n")
		if err := json.NewEncoder(os.Stdout).Encode(msg); err != nil {
			log.Printf("could not encode json: %v", err)
		}
	}
	lengthPadding := make([]byte, 4)
	binary.LittleEndian.PutUint32(lengthPadding, uint32(len(msgBytes)))
	if _, err = t.writer.Write(append(lengthPadding, msgBytes...)); err != nil {
		return err
	}
	return nil
}

func (t *pipeTransport) Close() error {
	select {
	case <-t.closed:
		return nil
	default:
		return t.onClose()
	}
}

func (t *pipeTransport) isClosed() bool {
	select {
	case <-t.closed:
		return true
	default:
		return false
	}
}

func newPipeTransport(driverCli string) (transport, error) {
	t := &pipeTransport{
		closed: make(chan struct{}, 1),
	}

	cmd := exec.Command(driverCli, "run-driver")
	cmd.SysProcAttr = defaultSysProcAttr
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("could not create stdin pipe: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("could not create stdout pipe: %w", err)
	}
	t.writer = stdin
	t.bufReader = bufio.NewReader(stdout)

	t.onClose = func() error {
		select {
		case <-t.closed:
		default:
			close(t.closed)
		}
		if err := t.writer.Close(); err != nil {
			return err
		}
		// playwright-cli will exit when its stdin is closed
		if err := cmd.Wait(); err != nil {
			return err
		}
		return nil
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("could not start driver: %w", err)
	}

	return t, nil
}
