// +build windows

package playwright_test

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mxschmitt/playwright-go"
)

type remoteServer struct {
	url string
	cmd *exec.Cmd
}

func newRemoteServer() *remoteServer {
	driver, err := playwright.NewDriver(&playwright.RunOptions{})
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	cmd := exec.Command(driver.DriverBinaryLocation, "launch-server", browserName)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("could not get stdout pipe: %v", err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
	scanner := bufio.NewReader(stdout)
	url, err := scanner.ReadString('\n')
	url = strings.TrimRight(url, "\n")
	if err != nil {
		log.Fatalf("could not read url: %v", err)
	}
	return &remoteServer{
		url: url,
		cmd: cmd,
	}
}

func (s *remoteServer) Close() {
	cmd := exec.Command("cmd", "/C", "taskkill", "/T", "/F", "/PID", strconv.Itoa(s.cmd.Process.Pid))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
