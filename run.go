package playwright

import (
	"fmt"
	"log"
	"os/exec"
)

func Run() (*Playwright, error) {
	cmd := exec.Command("./driver")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("could not get stdin pipe: %v", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("could not get stdout pipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("could not start driver: %v", err)
	}
	connection := newConnection(stdin, stdout, cmd.Process.Kill)
	go func() {
		if err := connection.Start(); err != nil {
			log.Printf("could not start connection: %v", err)
		}
	}()
	obj, err := connection.CallOnObjectWithKnownName("Playwright")
	if err != nil {
		return nil, fmt.Errorf("could not call object: %v", err)
	}
	return obj.(*Playwright), nil
}
