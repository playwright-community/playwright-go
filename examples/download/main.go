package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mxschmitt/playwright-go"
)

func exitIfErrorf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	startHttpServer()

	pw, err := playwright.Run()
	exitIfErrorf("could not launch playwright: %v", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	exitIfErrorf("could not launch Chromium: %v", err)
	context, err := browser.NewContext()
	exitIfErrorf("could not create context: %v", err)
	page, err := context.NewPage()
	exitIfErrorf("could not create page: %v", err)
	_, err = page.Goto("http://localhost:1234")
	exitIfErrorf("could not goto: %v", err)
	err = page.SetContent(`<a href="/download" download>download</a>`)
	exitIfErrorf("could not set content: %v", err)
	download, err := page.ExpectDownload(func() error {
		return page.Click("text=download")
	})
	exitIfErrorf("could not download: %v", err)
	fmt.Println(download.SuggestedFilename())
	err = browser.Close()
	exitIfErrorf("could not close browser: %v", err)
	err = pw.Stop()
	exitIfErrorf("could not stop Playwright: %v", err)
}

func startHttpServer() {
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("Content-Disposition", "attachment; filename=file.txt")
		if _, err := w.Write([]byte("foobar")); err != nil {
			log.Printf("could not write: %v", err)
		}
	})
	go func() {
		log.Fatal(http.ListenAndServe(":1234", nil))
	}()
}
