package playwright_test

import (
	"log"

	"github.com/mxschmitt/playwright-go"
)

func assertErrorToNil(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func ExampleRun() {
	pw, err := playwright.Run()
	assertErrorToNil("could not launch playwright: %v", err)
	browser, err := pw.Chromium.Launch()
	assertErrorToNil("could not launch Chromium: %v", err)
	context, err := browser.NewContext()
	assertErrorToNil("could not create context: %v", err)
	page, err := context.NewPage()
	assertErrorToNil("could not create page: %v", err)
	_, err = page.Goto("http://whatsmyuseragent.org/")
	assertErrorToNil("could not goto: %v", err)
	_, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	})
	assertErrorToNil("could not create screenshot: %v", err)
	assertErrorToNil("could not close browser: %v", browser.Close())
	assertErrorToNil("could not stop Playwright: %v", pw.Stop())
}
