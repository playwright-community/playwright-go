package playwright_test

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/playwright-community/playwright-go"
)

type remoteServer struct {
	url string
	cmd *exec.Cmd
}

func newRemoteServer() (*remoteServer, error) {
	driver, err := playwright.NewDriver(&playwright.RunOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not start Playwright: %v", err)
	}
	node := "node"
	if runtime.GOOS == "windows" {
		node = "node.exe"
	}
	cliJs := filepath.Join(driver.DriverDirectory, "package", "lib", "cli", "cli.js")
	cmd := exec.Command(filepath.Join(driver.DriverDirectory, node), cliJs, "launch-server", "--browser", browserName)
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("could not get stdout pipe: %v", err)
	}
	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("could not start server: %v", err)
	}
	scanner := bufio.NewReader(stdout)
	url, err := scanner.ReadString('\n')
	url = strings.TrimRight(url, "\n")
	if err != nil {
		return nil, fmt.Errorf("could not read url: %v", err)
	}
	return &remoteServer{
		url: url,
		cmd: cmd,
	}, nil
}

func (s *remoteServer) Close() {
	_ = s.cmd.Process.Kill()
	_ = s.cmd.Wait()
}
