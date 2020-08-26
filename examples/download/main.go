package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mxschmitt/playwright-golang"
)

func exitIfError(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	startHttpServer()

	pw, err := playwright.Run()
	exitIfError("could not launch playwright: %v", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	exitIfError("could not launch Chromium: %v", err)
	context, err := browser.NewContext()
	exitIfError("could not create context: %v", err)
	page, err := context.NewPage()
	exitIfError("could not create page: %v", err)
	err = page.Goto("http://localhost:1234")
	exitIfError("could not goto: %v", err)
	err = page.SetContent(`<a href="/download" download>download</a>`)
	exitIfError("could not set content: %v", err)
	downloadChan := make(chan *playwright.Download, 1)
	page.On("download", func(ev ...interface{}) {
		downloadChan <- ev[0].(*playwright.Download)
	})
	err = page.Click("text=download")
	exitIfError("could not click: %v", err)
	download := <-downloadChan
	fmt.Println(download.SuggestedFilename())
	err = browser.Close()
	exitIfError("could not close browser: %v", err)
	err = pw.Stop()
	exitIfError("could not stop Playwright: %v", err)
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
