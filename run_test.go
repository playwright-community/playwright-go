package playwright

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestDriverInstall(t *testing.T) {
	driverPath := t.TempDir()
	driver, err := NewDriver(&RunOptions{
		DriverDirectory: driverPath,
		Browsers:        []string{"firefox"},
		Verbose:         true})
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
