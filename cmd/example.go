package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mxschmitt/playwright-golang"
)

func exitIfError(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	playwright, err := playwright.Run()
	exitIfError("could not launch playwright: %v", err)
	browser, err := playwright.Chromium.Launch()
	exitIfError("could not launch Chromium: %v", err)
	context, err := browser.NewContext()
	exitIfError("could not create context: %v", err)
	page, err := context.NewPage()
	exitIfError("could not create page: %v", err)
	err = page.Goto("http://whatsmyuseragent.org/")
	exitIfError("could not goto: %v", err)
	content, err := page.Content()
	exitIfError("could not get content: %v", err)
	fmt.Println(content)
	screenshot, err := page.Screenshot()
	exitIfError("could not create screenshot: %v", err)
	err = ioutil.WriteFile("foo.png", screenshot, 0644)
	exitIfError("could not write file: %v", err)
	err = browser.Close()
	exitIfError("could not close browser: %v", err)
	err = playwright.Stop()
	exitIfError("could not stop Playwright: %v", err)
}
