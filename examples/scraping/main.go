package main

import (
	"fmt"
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
	assertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch()
	assertErrorToNilf("could not launch Chromium: %w", err)
	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)
	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)
	_, err = page.Goto("https://news.ycombinator.com")
	assertErrorToNilf("could not goto: %w", err)

	entries, err := page.EvaluateOnSelectorAll(".athing", `elements => [...elements].map(el => {
      const linkElement = el.children[2].children[0]
      return {
        title: linkElement.innerHTML,
      }
    })`)
	assertErrorToNilf("could not eval: %w", err)

	for i, entry := range entries.([]interface{}) {
		title := entry.(map[string]interface{})["title"]
		fmt.Println(i+1, title)
	}
	err = browser.Close()
	assertErrorToNilf("could not close browser: %w", err)
	err = pw.Stop()
	assertErrorToNilf("could not stop Playwright: %w", err)
}
