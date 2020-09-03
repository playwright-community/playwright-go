package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mxschmitt/playwright-go"
)

func exitIfErrorf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func worker(id int, jobs <-chan string, results chan<- bool, browser *playwright.Browser) {
	for url := range jobs {
		fmt.Println("starting", url)

		context, err := browser.NewContext(playwright.BrowserNewContextOptions{
			UserAgent: playwright.String("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
		})
		exitIfErrorf("could not create context: %v", err)

		page, err := context.NewPage()
		exitIfErrorf("could not create page: %v", err)

		_, err = page.Goto("http://"+url, playwright.PageGotoOptions{
			WaitUntil: playwright.String("networkidle"),
		})
		if err != nil {
			log.Printf("could not goto: %s: %v", url, err)
			context.Close()
			results <- true
			continue
		}
		cwd, err := os.Getwd()
		if err != nil {
			exitIfErrorf("could not get cwd %v", err)
		}
		_, err = page.Screenshot(playwright.PageScreenshotOptions{
			Path: playwright.String(filepath.Join(cwd, "out", strings.Replace(url, ".", "-", -1)+".png")),
		})
		exitIfErrorf("could not create screenshot: %v", err)
		fmt.Println("finish", url)
		context.Close()
		results <- true
	}
}

func main() {
	log.Println("Downloading Alexa top domains")
	topDomains, err := getAlexaTopDomains()
	exitIfErrorf("could not get alexa top domains: %v", err)
	log.Println("Downloaded Alexa top domains successfully")

	cwd, err := os.Getwd()
	if err != nil {
		exitIfErrorf("could not get cwd %v", err)
	}
	if err := os.Mkdir(filepath.Join(cwd, "out"), 0777); err != nil && !os.IsExist(err) {
		exitIfErrorf("could not create output directory %v", err)
	}

	pw, err := playwright.Run()
	exitIfErrorf("could not launch playwright: %v", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	exitIfErrorf("could not launch Chromium: %v", err)

	const numJobs = 30
	jobs := make(chan string, numJobs)
	results := make(chan bool, numJobs)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results, browser)
	}
	for _, url := range topDomains[:numJobs] {
		jobs <- url
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	err = browser.Close()
	exitIfErrorf("could not close browser: %v", err)
	err = pw.Stop()
	exitIfErrorf("could not stop Playwright: %v", err)
}

func getAlexaTopDomains() ([]string, error) {
	resp, err := http.Get("http://s3.amazonaws.com/alexa-static/top-1m.csv.zip")
	if err != nil {
		return nil, fmt.Errorf("could not get: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body: %v", err)
	}
	defer resp.Body.Close()
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, fmt.Errorf("could not create zip reader: %v", err)
	}
	alexaFile, err := zipReader.File[0].Open()
	if err != nil {
		return nil, fmt.Errorf("could not read alexa file: %v", err)
	}
	defer alexaFile.Close()
	reader := csv.NewReader(alexaFile)
	out := make([]string, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, fmt.Errorf("could not read csv: %v", err)
		}
		out = append(out, record[1])
	}
}
