package playwright_test

import (
	"log"
	"os"
	"testing"

	"github.com/playwright-community/playwright-go"
)

// global variables, can be used in any tests
var (
	pw          *playwright.Playwright
	browser     playwright.Browser
	context     playwright.BrowserContext
	page        playwright.Page
	expect      playwright.PlaywrightAssertions
	isChromium  bool
	isFirefox   bool
	isWebKit    bool
	browserName = getBrowserName()
	browserType playwright.BrowserType
)

// default context options for most tests
var DEFAULT_CONTEXT_OPTIONS = playwright.BrowserNewContextOptions{
	AcceptDownloads: playwright.Bool(true),
	HasTouch:        playwright.Bool(true),
}

// TestMain is used to setup and teardown the tests
func TestMain(m *testing.M) {
	BeforeAll()
	code := m.Run()
	AfterAll()
	os.Exit(code)
}

// BeforeAll prepares the environment, including
//   - start Playwright driver
//   - launch browser depends on BROWSER env
//   - init web-first assertions, alias as `expect`
func BeforeAll() {
	var err error
	pw, err = playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	if browserName == "chromium" || browserName == "" {
		browserType = pw.Chromium
	} else if browserName == "firefox" {
		browserType = pw.Firefox
	} else if browserName == "webkit" {
		browserType = pw.WebKit
	}
	// launch browser, headless or not depending on HEADFUL env
	browser, err = browserType.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(os.Getenv("HEADFUL") == ""),
	})
	if err != nil {
		log.Fatalf("could not launch: %v", err)
	}
	// init web-first assertions with 1s timeout instead of default 5s
	expect = playwright.NewPlaywrightAssertions(1000)
	isChromium = browserName == "chromium" || browserName == ""
	isFirefox = browserName == "firefox"
	isWebKit = browserName == "webkit"

	// for playwright-go tests
	server = newTestServer()
	utils = &testUtils{}
}

// AfterAll does cleanup, e.g. stop playwright driver
func AfterAll() {
	if server != nil {
		server.testServer.Close()
	}
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
}

// BeforeEach creates a new context and page for each test,
// so each test has isolated environment. Usage:
//
//	Func TestFoo(t *testing.T) {
//	  BeforeEach(t)
//	  defer AfterEach(t)
//	  // your test code
//	}
func BeforeEach(t *testing.T, contextOptions ...playwright.BrowserNewContextOptions) {
	t.Helper()
	if len(contextOptions) == 1 {
		newContextWithOptions(t, contextOptions[0])
		return
	}
	newContextWithOptions(t, DEFAULT_CONTEXT_OPTIONS)
}

// AfterEach closes the context and page after each test
func AfterEach(t *testing.T, closeContext ...bool) {
	t.Helper()
	if len(closeContext) == 0 {
		if err := context.Close(); err != nil {
			t.Errorf("could not close context: %v", err)
		}
	}
	server.AfterEach()
}

func getBrowserName() string {
	browserName, hasEnv := os.LookupEnv("BROWSER")
	if hasEnv {
		return browserName
	}
	return "chromium"
}

func newContextWithOptions(t *testing.T, contextOptions playwright.BrowserNewContextOptions) {
	t.Helper()
	var err error
	context, err = browser.NewContext(contextOptions)
	if err != nil {
		t.Fatalf("could not create new context: %v", err)
	}
	page, err = context.NewPage()
	if err != nil {
		t.Fatalf("could not create new page: %v", err)
	}
}
