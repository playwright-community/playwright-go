package playwright

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type playwrightDriver struct {
	driverName, driverFolder, driverPath, version string
}

func newDriver() (*playwrightDriver, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not get cwd: %w", err)
	}
	driverFolder := filepath.Join(cwd, ".ms-playwright")
	driverName := getDriverName()
	driverPath := filepath.Join(driverFolder, driverName)
	/**
	https://playwright.azureedge.net/builds/cli/playwright-cli-0.161.0-mac.zip
	https://playwright.azureedge.net/builds/cli/playwright-cli-0.161.0-linux.zip
	https://playwright.azureedge.net/builds/cli/playwright-cli-0.161.0-win32_x64.zip
	*/
	return &playwrightDriver{
		driverPath:   driverPath,
		driverFolder: driverFolder,
		driverName:   driverName,
		version:      "0.170.0-next.1608058598043",
	}, nil
}

func (d *playwrightDriver) isUpToDate() (bool, error) {
	if _, err := os.Stat(d.driverFolder); os.IsNotExist(err) {
		if err := os.Mkdir(d.driverFolder, 0777); err != nil {
			return false, fmt.Errorf("could not create driver folder: %w", err)
		}
	}
	if _, err := os.Stat(d.driverPath); os.IsNotExist(err) {
		return false, nil
	}
	output, err := exec.Command(d.driverPath, "--version").Output()
	if err != nil {
		return false, fmt.Errorf("could not run driver: %w", err)
	}
	if bytes.Contains(output, []byte(d.version)) {
		return true, nil
	}
	return false, nil
}

func (d *playwrightDriver) install() error {
	up2Date, err := d.isUpToDate()
	if err != nil {
		return fmt.Errorf("could not check if driver is up2date: %w", err)
	}
	if up2Date {
		return nil
	}

	log.Println("Downloading driver...")
	driverURL := d.getDriverURL()
	resp, err := http.Get(driverURL)
	if err != nil {
		return fmt.Errorf("could not download driver: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: got non 200 status code: %d (%s)", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return fmt.Errorf("could not read zip content: %w", err)
	}

	for _, zipFile := range zipReader.File {
		outFile, err := os.Create(filepath.Join(d.driverFolder, zipFile.Name))
		if err != nil {
			return fmt.Errorf("could not create driver: %w", err)
		}
		file, err := zipFile.Open()
		if err != nil {
			return fmt.Errorf("could not open zip file: %w", err)
		}
		if _, err = io.Copy(outFile, file); err != nil {
			return fmt.Errorf("could not copy response body to file: %w", err)
		}
		if err := outFile.Close(); err != nil {
			return fmt.Errorf("could not close file (driver): %w", err)
		}
		if err := file.Close(); err != nil {
			return fmt.Errorf("could not close file (zip file): %w", err)
		}
	}

	if runtime.GOOS != "windows" {
		stats, err := os.Stat(d.driverPath)
		if err != nil {
			return fmt.Errorf("could not stat driver: %w", err)
		}
		if err := os.Chmod(d.driverPath, stats.Mode()|0x40); err != nil {
			return fmt.Errorf("could not set permissions: %w", err)
		}
	}
	log.Println("Downloaded driver successfully")

	log.Println("Downloading browsers...")
	if err := installBrowsers(d.driverPath); err != nil {
		return fmt.Errorf("could not install browsers: %w", err)
	}
	log.Println("Downloaded browsers successfully")
	return nil
}

func (d *playwrightDriver) run() (*Connection, error) {
	cmd := exec.Command(d.driverPath, "run-driver")
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
	return newConnection(stdin, stdout, cmd.Process.Kill), nil
}

func installBrowsers(driverPath string) error {
	cmd := exec.Command(driverPath, "install")
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
	driver, err := newDriver()
	if err != nil {
		return fmt.Errorf("could not get driver instance: %w", err)
	}
	if err := driver.install(); err != nil {
		return fmt.Errorf("could not install driver: %w", err)
	}
	return nil
}

func Run() (*Playwright, error) {
	driver, err := newDriver()
	if err != nil {
		return nil, fmt.Errorf("could not get driver instance: %w", err)
	}
	if err := driver.install(); err != nil {
		return nil, fmt.Errorf("could not install driver: %w", err)
	}
	connection, err := driver.run()
	if err != nil {
		return nil, err
	}
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

func getDriverName() string {
	switch runtime.GOOS {
	case "windows":
		return "playwright-cli.exe"
	case "darwin":
		return "playwright-cli"
	case "linux":
		return "playwright-cli"
	}
	panic("Not supported OS!")
}

func (d *playwrightDriver) getDriverURL() string {
	platform := ""
	switch runtime.GOOS {
	case "windows":
		platform = "win32_x64"
	case "darwin":
		platform = "mac"
	case "linux":
		platform = "linux"
	}
	return fmt.Sprintf("https://playwright.azureedge.net/builds/cli/next/playwright-cli-%s-%s.zip", d.version, platform)
}
