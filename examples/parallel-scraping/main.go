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

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func worker(id int, jobs chan job, results chan<- bool, browser *playwright.Browser) {
	for jobPayload := range jobs {
		if jobPayload.Try > 3 {
			log.Printf("Stopped with domain %s", jobPayload.URL)
			results <- true
			continue
		}
		fmt.Printf("starting (%d): %s\n", jobPayload.Try, jobPayload.URL)

		context, err := browser.NewContext(playwright.BrowserNewContextOptions{
			UserAgent: playwright.String("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
		})
		assertErrorToNilf("could not create context: %w", err)

		page, err := context.NewPage()
		assertErrorToNilf("could not create page: %w", err)

		_, err = page.Goto("http://"+jobPayload.URL, playwright.PageGotoOptions{
			WaitUntil: playwright.String("networkidle"),
		})
		if err != nil {
			log.Printf("could not goto: %s: %v", jobPayload.URL, err)
			context.Close()
			jobs <- job{
				URL: jobPayload.URL,
				Try: jobPayload.Try + 1,
			}
			continue
		}
		cwd, err := os.Getwd()
		if err != nil {
			assertErrorToNilf("could not get cwd %w", err)
		}
		_, err = page.Screenshot(playwright.PageScreenshotOptions{
			Path: playwright.String(filepath.Join(cwd, "out", strings.Replace(jobPayload.URL, ".", "-", -1)+".png")),
		})
		assertErrorToNilf("could not create screenshot: %w", err)
		fmt.Println("finish", jobPayload.URL)
		context.Close()
		results <- true
	}
}

type job struct {
	URL string
	Try int
}

func main() {
	log.Println("Downloading Alexa top domains")
	topDomains, err := getAlexaTopDomains()
	assertErrorToNilf("could not get alexa top domains: %w", err)
	log.Println("Downloaded Alexa top domains successfully")

	cwd, err := os.Getwd()
	if err != nil {
		assertErrorToNilf("could not get cwd %w", err)
	}
	if err := os.Mkdir(filepath.Join(cwd, "out"), 0777); err != nil && !os.IsExist(err) {
		assertErrorToNilf("could not create output directory %w", err)
	}

	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	const numJobs = 30
	jobs := make(chan job, numJobs)
	results := make(chan bool, numJobs)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results, browser)
	}
	for _, url := range topDomains[:numJobs] {
		jobs <- job{
			URL: url,
		}
	}
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	assertErrorToNilf("could not close browser: %w", browser.Close())
	assertErrorToNilf("could not stop Playwright: %w", pw.Stop())
}

func getAlexaTopDomains() ([]string, error) {
	resp, err := http.Get("http://s3.amazonaws.com/alexa-static/top-1m.csv.zip")
	if err != nil {
		return nil, fmt.Errorf("could not get: %w", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body: %w", err)
	}
	defer resp.Body.Close()
	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return nil, fmt.Errorf("could not create zip reader: %w", err)
	}
	alexaFile, err := zipReader.File[0].Open()
	if err != nil {
		return nil, fmt.Errorf("could not read alexa file: %w", err)
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
			return nil, fmt.Errorf("could not read csv: %w", err)
		}
		out = append(out, record[1])
	}
}
