package playwright

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	playwrightCliVersion = "1.25.2"
	baseURL              = "https://playwright.azureedge.net/builds/driver"
)

type PlaywrightDriver struct {
	DriverDirectory, DriverBinaryLocation, Version string
	options                                        *RunOptions
}

func NewDriver(options *RunOptions) (*PlaywrightDriver, error) {
	baseDriverDirectory := options.DriverDirectory
	if baseDriverDirectory == "" {
		var err error
		baseDriverDirectory, err = getDefaultCacheDirectory()
		if err != nil {
			return nil, fmt.Errorf("could not get default cache directory: %v", err)
		}
	}
	driverDirectory := filepath.Join(baseDriverDirectory, "ms-playwright-go", playwrightCliVersion)
	driverBinaryLocation := filepath.Join(driverDirectory, getDriverName())
	return &PlaywrightDriver{
		options:              options,
		DriverBinaryLocation: driverBinaryLocation,
		DriverDirectory:      driverDirectory,
		Version:              playwrightCliVersion,
	}, nil
}

func getDefaultCacheDirectory() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %v", err)
	}
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(userHomeDir, "AppData", "Local"), nil
	case "darwin":
		return filepath.Join(userHomeDir, "Library", "Caches"), nil
	case "linux":
		return filepath.Join(userHomeDir, ".cache"), nil
	}
	return "", errors.New("could not determine cache directory")
}

func (d *PlaywrightDriver) isUpToDateDriver() (bool, error) {
	if _, err := os.Stat(d.DriverDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(d.DriverDirectory, 0777); err != nil {
			return false, fmt.Errorf("could not create driver directory: %w", err)
		}
	}
	if _, err := os.Stat(d.DriverBinaryLocation); os.IsNotExist(err) {
		return false, nil
	}
	cmd := exec.Command(d.DriverBinaryLocation, "--version")
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("could not run driver: %w", err)
	}
	if bytes.Contains(output, []byte(d.Version)) {
		return true, nil
	}
	return false, nil
}

func (d *PlaywrightDriver) install() error {
	if err := d.DownloadDriver(); err != nil {
		return fmt.Errorf("could not install driver: %w", err)
	}
	if d.options.SkipInstallBrowsers {
		return nil
	}
	if d.options.Verbose {
		log.Println("Downloading browsers...")
	}
	if err := d.installBrowsers(d.DriverBinaryLocation); err != nil {
		return fmt.Errorf("could not install browsers: %w", err)
	}
	if d.options.Verbose {
		log.Println("Downloaded browsers successfully")
	}
	return nil
}
func (d *PlaywrightDriver) DownloadDriver() error {
	up2Date, err := d.isUpToDateDriver()
	if err != nil {
		return fmt.Errorf("could not check if driver is up2date: %w", err)
	}
	if up2Date {
		return nil
	}

	if d.options.Verbose {
		log.Printf("Downloading driver to %s", d.DriverDirectory)
	}
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
		zipFileDiskPath := filepath.Join(d.DriverDirectory, zipFile.Name)
		if zipFile.FileInfo().IsDir() {
			if err := os.MkdirAll(zipFileDiskPath, os.ModePerm); err != nil {
				return fmt.Errorf("could not create directory: %w", err)
			}
			continue
		}

		outFile, err := os.Create(zipFileDiskPath)
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
		if zipFile.Mode().Perm()&0100 != 0 && runtime.GOOS != "windows" {
			if err := makeFileExecutable(zipFileDiskPath); err != nil {
				return fmt.Errorf("could not make executable: %w", err)
			}
		}
	}

	if d.options.Verbose {
		log.Println("Downloaded driver successfully")
	}
	return nil
}

func (d *PlaywrightDriver) run() (*connection, error) {
	cmd := exec.Command(d.DriverBinaryLocation, "run-driver")
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
	transport := newPipeTransport(stdin, stdout)
	go func() {
		if err := transport.Start(); err != nil {
			log.Fatal(err)
		}
	}()
	connection := newConnection(func() error {
		if err := stdin.Close(); err != nil {
			return fmt.Errorf("could not close stdin: %v", err)
		}
		if err := stdout.Close(); err != nil {
			return fmt.Errorf("could not close stdout: %v", err)
		}
		if err := cmd.Process.Kill(); err != nil {
			return fmt.Errorf("could not kill process: %v", err)
		}
		if _, err := cmd.Process.Wait(); err != nil {
			return fmt.Errorf("could not wait for process: %v", err)
		}
		return nil
	})
	connection.onmessage = transport.Send
	transport.onmessage = connection.Dispatch
	return connection, nil
}

func (d *PlaywrightDriver) installBrowsers(driverPath string) error {
	additionalArgs := []string{"install"}
	if d.options.Browsers != nil {
		additionalArgs = append(additionalArgs, d.options.Browsers...)
	}
	cmd := exec.Command(driverPath, additionalArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("could not install browsers: %w", err)
	}
	return nil
}

// RunOptions are custom options to run the driver
type RunOptions struct {
	DriverDirectory     string
	SkipInstallBrowsers bool
	Browsers            []string
	Verbose             bool
}

// Install does download the driver and the browsers. If not called manually
// before playwright.Run() it will get executed there and might take a few seconds
// to download the Playwright suite.
func Install(options ...*RunOptions) error {
	driver, err := NewDriver(transformRunOptions(options))
	if err != nil {
		return fmt.Errorf("could not get driver instance: %w", err)
	}
	if err := driver.install(); err != nil {
		return fmt.Errorf("could not install driver: %w", err)
	}
	return nil
}

// Run starts a Playwright instance
func Run(options ...*RunOptions) (*Playwright, error) {
	driver, err := NewDriver(transformRunOptions(options))
	if err != nil {
		return nil, fmt.Errorf("could not get driver instance: %w", err)
	}
	connection, err := driver.run()
	if err != nil {
		return nil, err
	}
	playwright := connection.Start()
	return playwright, nil
}

func transformRunOptions(options []*RunOptions) *RunOptions {
	if len(options) == 1 {
		return options[0]
	}
	return &RunOptions{
		Verbose: true,
	}
}

func getDriverName() string {
	switch runtime.GOOS {
	case "windows":
		return "playwright.cmd"
	case "darwin":
		fallthrough
	case "linux":
		return "playwright.sh"
	}
	panic("Not supported OS!")
}

func (d *PlaywrightDriver) getDriverURL() string {
	platform := ""
	switch runtime.GOOS {
	case "windows":
		platform = "win32_x64"
	case "darwin":
		if runtime.GOARCH == "arm64" {
			platform = "mac-arm64"
		} else {
			platform = "mac"
		}
	case "linux":
		if runtime.GOARCH == "arm64" {
			platform = "linux-arm64"
		} else {
			platform = "linux"
		}
	}

	if d.isReleaseVersion() {
		return fmt.Sprintf("%s/playwright-%s-%s.zip", baseURL, d.Version, platform)
	}

	return fmt.Sprintf("%s/next/playwright-%s-%s.zip", baseURL, d.Version, platform)
}

// isReleaseVersion checks if the version is not a beta or alpha release
// this helps to determine the url from where to download the driver
func (d *PlaywrightDriver) isReleaseVersion() bool {
	return !strings.Contains(d.Version, "beta") && !strings.Contains(d.Version, "alpha")
}

func makeFileExecutable(path string) error {
	stats, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("could not stat driver: %w", err)
	}
	if err := os.Chmod(path, stats.Mode()|0x40); err != nil {
		return fmt.Errorf("could not set permissions: %w", err)
	}
	return nil
}
