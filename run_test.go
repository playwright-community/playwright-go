package playwright

import (
	"os"
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
