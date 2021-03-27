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

const playwrightCliVersion = "1.10.0"

type playwrightDriver struct {
	driverDirectory, driverBinaryLocation, version string
	options                                        *RunOptions
}

func newDriver(options *RunOptions) (*playwrightDriver, error) {
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
	return &playwrightDriver{
		options:              options,
		driverBinaryLocation: driverBinaryLocation,
		driverDirectory:      driverDirectory,
		version:              playwrightCliVersion,
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

func (d *playwrightDriver) isUpToDateDriver() (bool, error) {
	if _, err := os.Stat(d.driverDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(d.driverDirectory, 0777); err != nil {
			return false, fmt.Errorf("could not create driver directory: %w", err)
		}
	}
	if _, err := os.Stat(d.driverBinaryLocation); os.IsNotExist(err) {
		return false, nil
	}
	cmd := exec.Command(d.driverBinaryLocation, "--version")
	cmd.Env = d.getDriverEnviron()
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("could not run driver: %w", err)
	}
	if bytes.Contains(output, []byte(d.version)) {
		return true, nil
	}
	return false, nil
}

func (d *playwrightDriver) install() error {
	if err := d.installDriver(); err != nil {
		return fmt.Errorf("could not install driver: %w", err)
	}
	if d.options.SkipInstallBrowsers {
		return nil
	}
	log.Println("Downloading browsers...")
	if err := d.installBrowsers(d.driverBinaryLocation); err != nil {
		return fmt.Errorf("could not install browsers: %w", err)
	}
	log.Println("Downloaded browsers successfully")
	return nil
}
func (d *playwrightDriver) installDriver() error {
	up2Date, err := d.isUpToDateDriver()
	if err != nil {
		return fmt.Errorf("could not check if driver is up2date: %w", err)
	}
	if up2Date {
		return nil
	}

	log.Printf("Downloading driver to %s", d.driverDirectory)
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
		zipFileDiskPath := filepath.Join(d.driverDirectory, zipFile.Name)
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

	log.Println("Downloaded driver successfully")
	return nil
}

func (d *playwrightDriver) run() (*connection, error) {
	cmd := exec.Command(d.driverBinaryLocation, "run-driver")
	cmd.Env = d.getDriverEnviron()
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

func (d *playwrightDriver) installBrowsers(driverPath string) error {
	additionalArgs := []string{"install"}
	if d.options.Browsers != nil {
		additionalArgs = append(additionalArgs, d.options.Browsers...)
	}
	cmd := exec.Command(driverPath, additionalArgs...)
	cmd.Env = d.getDriverEnviron()
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

// RunOptions are custom options to run the driver
type RunOptions struct {
	DriverDirectory     string
	SkipInstallBrowsers bool
	Browsers            []string
}

// Install does download the driver and the browsers. If not called manually
// before playwright.Run() it will get executed there and might take a few seconds
// to download the Playwright suite.
func Install(options ...*RunOptions) error {
	driver, err := newDriver(transformRunOptions(options))
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
	driver, err := newDriver(transformRunOptions(options))
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

func transformRunOptions(options []*RunOptions) *RunOptions {
	if len(options) == 1 {
		return options[0]
	}
	return &RunOptions{}
}

func getDriverName() string {
	switch runtime.GOOS {
	case "windows":
		return "playwright.cmd"
	case "darwin":
		return "playwright.sh"
	case "linux":
		return "playwright.sh"
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
	optionalSubDirectory := ""
	if strings.Contains(d.version, "next") {
		optionalSubDirectory = "/next"
	}
	return fmt.Sprintf("https://playwright.azureedge.net/builds/driver%s/playwright-%s-%s.zip", optionalSubDirectory, d.version, platform)
}

func (d *playwrightDriver) getDriverEnviron() []string {
	environ := os.Environ()
	unset := func(key string) {
		for i := range environ {
			if strings.HasPrefix((environ)[i], key+"=") {
				(environ)[i] = (environ)[len(environ)-1]
				environ = (environ)[:len(environ)-1]
				break
			}
		}
	}
	unset("NODE_OPTIONS")
	return environ
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
