package playwright

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func getDriverURL() (string, string) {
	const baseURL = "https://storage.googleapis.com/mxschmitt-public-files/"
	const version = "playwright-driver-1599218722159"
	driverName := ""
	switch runtime.GOOS {
	case "windows":
		driverName = "playwright-driver-win.exe"
	case "darwin":
		driverName = "playwright-driver-macos"
	case "linux":
		driverName = "playwright-driver-linux"
	}
	return fmt.Sprintf("%s%s/%s", baseURL, version, driverName), driverName
}

func installPlaywright() (string, error) {
	driverURL, driverName := getDriverURL()
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("could not get cwd: %w", err)
	}
	driverFolder := filepath.Join(cwd, ".ms-playwright")
	if _, err = os.Stat(driverFolder); os.IsNotExist(err) {
		if err := os.Mkdir(driverFolder, 0777); err != nil {
			return "", fmt.Errorf("could not create driver folder :%w", err)
		}
	}
	driverPath := filepath.Join(driverFolder, driverName)
	if _, err := os.Stat(driverPath); err == nil {
		return driverPath, nil
	}
	log.Println("Downloading driver...")
	resp, err := http.Get(driverURL)
	if err != nil {
		return "", fmt.Errorf("could not download driver: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: got non 2xx status code: %d (%s)", resp.StatusCode, resp.Status)
	}
	outFile, err := os.Create(driverPath)
	if err != nil {
		return "", fmt.Errorf("could not create driver: %w", err)
	}
	if _, err = io.Copy(outFile, resp.Body); err != nil {
		return "", fmt.Errorf("could not copy response body to file: %w", err)
	}
	if err := outFile.Close(); err != nil {
		return "", fmt.Errorf("could not close file (driver): %w", err)
	}

	if runtime.GOOS != "windows" {
		stats, err := os.Stat(driverPath)
		if err != nil {
			return "", fmt.Errorf("could not stat driver: %w", err)
		}
		if err := os.Chmod(driverPath, stats.Mode()|0x40); err != nil {
			return "", fmt.Errorf("could not set permissions: %w", err)
		}
	}
	log.Println("Downloaded driver successfully")

	log.Println("Downloading browsers...")
	if err := installBrowsers(driverPath); err != nil {
		return "", fmt.Errorf("could not install browsers: %w", err)
	}
	log.Println("Downloaded browsers successfully")
	return driverPath, nil
}

func installBrowsers(driverPath string) error {
	cmd := exec.Command(driverPath, "--install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("could not start driver: %w", err)
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

// Install does download the driver and the browsers. If not called manually
// before playwright.Run() it will get executed there and might take a few seconds
// to download the Playwright suite.
func Install() error {
	_, err := installPlaywright()
	if err != nil {
		return fmt.Errorf("could not install driver: %w", err)
	}
	return nil
}

func Run() (*Playwright, error) {
	driverPath, err := installPlaywright()
	if err != nil {
		return nil, fmt.Errorf("could not install driver: %w", err)
	}

	cmd := exec.Command(driverPath, "--run")
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("could not get stdin pipe: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("could not get stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("could not start driver: %w", err)
	}
	connection := newConnection(stdin, stdout, cmd.Process.Kill)
	go func() {
		if err := connection.Start(); err != nil {
			log.Fatalf("could not start connection: %v", err)
		}
	}()
	obj, err := connection.CallOnObjectWithKnownName("Playwright")
	if err != nil {
		return nil, fmt.Errorf("could not call object: %w", err)
	}
	return obj.(*Playwright), nil
}
