package main

import (
	"fmt"
	"log"

	"github.com/mxschmitt/playwright-go"
)

func exitIfErrorf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	pw, err := playwright.Run()
	exitIfErrorf("could not launch playwright: %v", err)
	browser, err := pw.Chromium.Launch()
	exitIfErrorf("could not launch Chromium: %v", err)
	context, err := browser.NewContext()
	exitIfErrorf("could not create context: %v", err)
	page, err := context.NewPage()
	exitIfErrorf("could not create page: %v", err)
	err = page.Goto("https://news.ycombinator.com")
	exitIfErrorf("could not goto: %v", err)

	entries, err := page.EvaluateOnSelectorAll(".athing", `elements => [...elements].map(el => {
      const linkElement = el.children[2].children[0]
      return {
        title: linkElement.innerHTML,
      }
    })`)
	exitIfErrorf("could not eval: %v", err)

	for i, entry := range entries.([]interface{}) {
		title := entry.(map[string]interface{})["title"]
		fmt.Println(i+1, title)
	}
	err = browser.Close()
	exitIfErrorf("could not close browser: %v", err)
	err = pw.Stop()
	exitIfErrorf("could not stop Playwright: %v", err)
}
