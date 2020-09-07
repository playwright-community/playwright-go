package main

import (
	"log"

	"github.com/mxschmitt/playwright-go"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not launch playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch Chromium: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	_, err = page.Goto("http://whatsmyuseragent.org/", playwright.PageGotoOptions{
		WaitUntil: playwright.String("networkidle"),
	})
	if err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	_, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	})
	if err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
	err = browser.Close()
	if err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	err = pw.Stop()
	if err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
