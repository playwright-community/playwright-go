package main

import (
	"log"

	"github.com/mxschmitt/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %v", err)
	browser, err := pw.Chromium.Launch()
	assertErrorToNilf("could not launch Chromium: %v", err)
	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %v", err)
	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %v", err)
	_, err = page.Goto("http://whatsmyuseragent.org/")
	assertErrorToNilf("could not goto: %v", err)
	_, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	})
	assertErrorToNilf("could not create screenshot: %v", err)
	err = browser.Close()
	assertErrorToNilf("could not close browser: %v", err)
	err = pw.Stop()
	assertErrorToNilf("could not stop Playwright: %v", err)
}
