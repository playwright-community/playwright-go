package main

import (
	"fmt"
	"log"

	"github.com/mxschmitt/playwright-go"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	_, err = page.Goto("https://news.ycombinator.com")
	if err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	entries, err := page.QuerySelectorAll(".athing")
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	for i, entry := range entries {
		titleElement, err := entry.QuerySelector("td.title > a")
		if err != nil {
			log.Fatalf("could not get title element: %v", err)
		}
		title, err := titleElement.TextContent()
		if err != nil {
			log.Fatalf("could not get title: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, title)
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
