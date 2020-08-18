package playwright

import (
	"crypto/sha1"
	"encoding/hex"
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
	const version = "playwright-driver-1597776060158"
	driverName := ""
	switch runtime.GOOS {
	case "windows":
		driverName = "playwright-driver-win.exe"
		break
	case "darwin":
		driverName = "playwright-driver-macos"
		break
	case "linux":
		driverName = "playwright-driver-linux"
		break
	}
	hash := sha1.New()
	hash.Write([]byte(version))
	return fmt.Sprintf("%s%s/%s", baseURL, version, driverName), fmt.Sprintf("%s-%s", driverName, hex.EncodeToString(hash.Sum(nil))[:5])
}

func installPlaywright() (string, error) {
	driverURL, driverName := getDriverURL()
	driverPath := filepath.Join(os.TempDir(), driverName)
	_, err := os.Stat(driverPath)
	if err == nil {
		return driverPath, nil
	}
	if !os.IsNotExist(err) {
		return driverName, err
	}
	log.Println("Downloading driver...")
	resp, err := http.Get(driverURL)
	if err != nil {
		return "", fmt.Errorf("could not download driver: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: got non 2xx status code: %d (%s)", resp.StatusCode, resp.Status)
	}
	outFile, err := os.Create(driverPath)
	if err != nil {
		return "", fmt.Errorf("could not create driver: %v", err)
	}
	if _, err = io.Copy(outFile, resp.Body); err != nil {
		return "", fmt.Errorf("could not copy response body to file: %v", err)
	}
	if err := outFile.Close(); err != nil {
		return "", fmt.Errorf("could not close file (driver): %v", err)
	}

	if runtime.GOOS != "windows" {
		stats, err := os.Stat(driverPath)
		if err != nil {
			return "", fmt.Errorf("could not stat driver: %v", err)
		}
		if err := os.Chmod(driverPath, stats.Mode()|0x40); err != nil {
			return "", fmt.Errorf("could not set permissions: %v", err)
		}
	}
	log.Println("Downloaded driver successfully")

	log.Println("Downloading browsers...")
	if err := installBrowsers(driverPath); err != nil {
		return "", fmt.Errorf("could not install browsers: %v", err)
	}
	log.Println("Downloaded browsers successfully")
	return driverPath, nil
}

func installBrowsers(driverPath string) error {
	cmd := exec.Command(driverPath, "--install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("could not start driver: %v", err)
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
		return fmt.Errorf("could not install driver: %v", err)
	}
	return nil
}

func Run() (*Playwright, error) {
	driverPath, err := installPlaywright()
	if err != nil {
		return nil, fmt.Errorf("could not install driver: %v", err)
	}

	cmd := exec.Command(driverPath, "--run")
	cmd.Stderr = os.Stderr
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
