package main

import (
	"fmt"
	"log"

	"github.com/mxschmitt/playwright-golang"
)

func exitIfError(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	pw, err := playwright.Run()
	exitIfError("could not launch playwright: %v", err)
	browser, err := pw.Chromium.Launch()
	exitIfError("could not launch Chromium: %v", err)
	context, err := browser.NewContext()
	exitIfError("could not create context: %v", err)
	page, err := context.NewPage()
	exitIfError("could not create page: %v", err)
	err = page.Goto("https://news.ycombinator.com")
	exitIfError("could not goto: %v", err)

	entries, err := page.EvaluateOnSelectorAll(".athing", `elements => [...elements].map(el => {
      const linkElement = el.children[2].children[0]
      return {
        title: linkElement.innerHTML,
      }
    })`)
	exitIfError("could not eval: %v", err)

	for i, entry := range entries.([]interface{}) {
		title := entry.(map[string]interface{})["title"]
		fmt.Println(i+1, title)
	}
	err = browser.Close()
	exitIfError("could not close browser: %v", err)
	err = pw.Stop()
	exitIfError("could not stop Playwright: %v", err)
}
