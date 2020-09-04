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
	assertErrorToNilf("could not launch playwright: %v", err)
	browser, err := pw.Chromium.Launch()
	assertErrorToNilf("could not launch Chromium: %v", err)
	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %v", err)
	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %v", err)
	_, err = page.Goto("https://news.ycombinator.com")
	assertErrorToNilf("could not goto: %v", err)

	entries, err := page.EvaluateOnSelectorAll(".athing", `elements => [...elements].map(el => {
      const linkElement = el.children[2].children[0]
      return {
        title: linkElement.innerHTML,
      }
    })`)
	assertErrorToNilf("could not eval: %v", err)

	for i, entry := range entries.([]interface{}) {
		title := entry.(map[string]interface{})["title"]
		fmt.Println(i+1, title)
	}
	err = browser.Close()
	assertErrorToNilf("could not close browser: %v", err)
	err = pw.Stop()
	assertErrorToNilf("could not stop Playwright: %v", err)
}
