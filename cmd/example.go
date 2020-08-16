package main

import (
	"log"

	"github.com/mxschmitt/playwright-golang"
)

func main() {
	playwright, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not launch playwright: %v", err)
	}
	browser, err := playwright.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch Chromium: %v", err)
	}
	context, err := browser.NewContext()
	if err != nil {
		log.Fatalf("could not create context: %v", err)
	}
	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if err = page.Goto("http://whatsmyuseragent.org/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	if err = page.Screenshot("example.png"); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = playwright.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
