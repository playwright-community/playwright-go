package playwright

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/mitchellh/go-ps"
	"github.com/stretchr/testify/require"
)

func TestDriverInstall(t *testing.T) {
	driverPath := t.TempDir()
	driver, err := NewDriver(&RunOptions{
		DriverDirectory: driverPath,
		Browsers:        []string{getBrowserName()},
		Verbose:         true,
	})
	if err != nil {
		t.Fatalf("could not start driver: %v", err)
	}
	browserPath := t.TempDir()
	err = os.Setenv("PLAYWRIGHT_BROWSERS_PATH", browserPath)
	if err != nil {
		t.Fatalf("could not set PLAYWRIGHT_BROWSERS_PATH: %v", err)
	}
	defer os.Unsetenv("PLAYWRIGHT_BROWSERS_PATH")
	err = driver.Install()
	if err != nil {
		t.Fatalf("could not install driver: %v", err)
	}
	err = driver.Uninstall()
	if err != nil {
		t.Fatalf("could not uninstall driver: %v", err)
	}
}

func TestDriverDownloadHostEnv(t *testing.T) {
	driverPath := t.TempDir()
	driver, err := NewDriver(&RunOptions{
		DriverDirectory:     driverPath,
		SkipInstallBrowsers: true,
	})
	if err != nil {
		t.Fatalf("could not start driver: %v", err)
	}
	uri := ""
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri = r.URL.String()
		w.WriteHeader(404)
	}))
	defer ts.Close()

	err = os.Setenv("PLAYWRIGHT_DOWNLOAD_HOST", ts.URL)
	if err != nil {
		t.Fatalf("could not set PLAYWRIGHT_DOWNLOAD_HOST: %v", err)
	}
	defer os.Unsetenv("PLAYWRIGHT_DOWNLOAD_HOST")
	err = driver.Install()
	if err == nil || !strings.Contains(err.Error(), "404 Not Found") || !strings.Contains(uri, "/builds/driver") {
		t.Fatalf("PLAYWRIGHT_DOWNLOAD_HOST do not work: %v", err)
	}
}

func TestShouldNotHangWhenPlaywrightUnexpectedExit(t *testing.T) {
	if getBrowserName() != "chromium" {
		t.Skip("chromium only")
		return
	}

	pw, err := Run()
	require.NoError(t, err)
	defer func() {
		_ = pw.Stop()
	}()
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	context, err := browser.NewContext()
	require.NoError(t, err)

	err = killPlaywrightProcess()
	require.NoError(t, err)

	_, err = context.NewPage()
	require.Error(t, err)
}

// find and kill playwright process
func killPlaywrightProcess() error {
	all, err := ps.Processes()
	if err != nil {
		return err
	}
	for _, process := range all {
		if process.Executable() == "node" || process.Executable() == "node.exe" {
			if process.PPid() == os.Getpid() {
				if err := killProcessByPid(process.Pid()); err != nil {
					return err
				}
				return nil
			}
		}
	}
	return fmt.Errorf("playwright process not found")
}

func killProcessByPid(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	if err := process.Kill(); err != nil {
		return err
	}
	return nil
}

func getBrowserName() string {
	browserName, hasEnv := os.LookupEnv("BROWSER")
	if hasEnv {
		return browserName
	}
	return "chromium"
}
