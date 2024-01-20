# 🎭 [Playwright](https://github.com/microsoft/playwright#readme) for <img src="https://user-images.githubusercontent.com/17984549/91302719-343a1d80-e7a7-11ea-8d6a-9448ef598420.png" height="35" />

## Looking for maintainers and see [here](https://github.com/playwright-community/playwright-go/issues/122). Thanks!

[![PkgGoDev](https://pkg.go.dev/badge/github.com/playwright-community/playwright-go)](https://pkg.go.dev/github.com/playwright-community/playwright-go)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](http://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/playwright-community/playwright-go)](https://goreportcard.com/report/github.com/playwright-community/playwright-go) ![Build Status](https://github.com/playwright-community/playwright-go/workflows/Go/badge.svg)
[![Join Slack](https://img.shields.io/badge/join-slack-infomational)](https://aka.ms/playwright-slack) [![Coverage Status](https://coveralls.io/repos/github/playwright-community/playwright-go/badge.svg?branch=main)](https://coveralls.io/github/playwright-community/playwright-go?branch=main) <!-- GEN:chromium-version-badge -->[![Chromium version](https://img.shields.io/badge/chromium-121.0.6167.57-blue.svg?logo=google-chrome)](https://www.chromium.org/Home)<!-- GEN:stop --> <!-- GEN:firefox-version-badge -->[![Firefox version](https://img.shields.io/badge/firefox-121.0-blue.svg?logo=mozilla-firefox)](https://www.mozilla.org/en-US/firefox/new/)<!-- GEN:stop --> <!-- GEN:webkit-version-badge -->[![WebKit version](https://img.shields.io/badge/webkit-17.4-blue.svg?logo=safari)](https://webkit.org/)<!-- GEN:stop -->

[API reference](https://playwright.dev/docs/api/class-playwright) | [Example recipes](https://github.com/playwright-community/playwright-go/tree/main/examples)

Playwright is a Go library to automate [Chromium](https://www.chromium.org/Home), [Firefox](https://www.mozilla.org/en-US/firefox/new/) and [WebKit](https://webkit.org/) with a single API. Playwright is built to enable cross-browser web automation that is **ever-green**, **capable**, **reliable** and **fast**.

|          | Linux | macOS | Windows |
|   :---   | :---: | :---: | :---:   |
| Chromium <!-- GEN:chromium-version -->121.0.6167.57<!-- GEN:stop --> | ✅ | ✅ | ✅ |
| WebKit <!-- GEN:webkit-version -->17.4<!-- GEN:stop --> | ✅ | ✅ | ✅ |
| Firefox <!-- GEN:firefox-version -->121.0<!-- GEN:stop --> | ✅ | ✅ | ✅ |

Headless execution is supported for all the browsers on all platforms.

## Installation

```txt
go get -u github.com/playwright-community/playwright-go
```

Install the browsers and OS dependencies:

```bash
go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps
# Or
go install github.com/playwright-community/playwright-go/cmd/playwright@latest
playwright install --with-deps
```

Alternatively you can do it inside your program via which downloads the driver and browsers:

```go
err := playwright.Install()
```

## Capabilities

Playwright is built to automate the broad and growing set of web browser capabilities used by Single Page Apps and Progressive Web Apps.

* Scenarios that span multiple page, domains and iframes
* Auto-wait for elements to be ready before executing actions (like click, fill)
* Intercept network activity for stubbing and mocking network requests
* Emulate mobile devices, geolocation, permissions
* Support for web components via shadow-piercing selectors
* Native input events for mouse and keyboard
* Upload and download files

## Example

The following example crawls the current top voted items from [Hacker News](https://news.ycombinator.com).

```go

package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
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
	if _, err = page.Goto("https://news.ycombinator.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	entries, err := page.Locator(".athing").All()
	if err != nil {
		log.Fatalf("could not get entries: %v", err)
	}
	for i, entry := range entries {
		title, err := entry.Locator("td.title > span > a").TextContent()
		if err != nil {
			log.Fatalf("could not get text content: %v", err)
		}
		fmt.Printf("%d: %s\n", i+1, title)
	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
```

## More examples

* [Downloading files](./examples/download/main.go)
* [End-To-End testing a website](./examples/end-to-end-testing/main.go)
* [Executing JavaScript in the browser](./examples/javascript/main.go)
* [Emulate mobile and geolocation](./examples/mobile-and-geolocation/main.go)
* [Parallel scraping using a WaitGroup](./examples/parallel-scraping/main.go)
* [Rendering a PDF of a website](./examples/pdf/main.go)
* [Scraping HackerNews](./examples/scraping/main.go)
* [Take a screenshot](./examples/screenshot/main.go)
* [Record a video](./examples/video/main.go)
* [Monitor network activity](./examples/network-monitoring/main.go)

## How does it work?

Playwright is a Node.js library which uses:

* Chrome DevTools Protocol to communicate with Chromium
* Patched Firefox to communicate with Firefox
* Patched WebKit to communicate with WebKit

These patches are based on the original sources of the browsers and don't modify the browser behaviour so the browsers are basically the same (see [here](https://github.com/microsoft/playwright/tree/main/browser_patches)) as you see them in the wild. The support for different programming languages is based on exposing a RPC server in the Node.js land which can be used to allow other languages to use Playwright without implementing all the custom logic:

* [Playwright for Python](https://github.com/microsoft/playwright-python)
* [Playwright for .NET](https://github.com/microsoft/playwright-sharp)
* [Playwright for Java](https://github.com/microsoft/playwright-java)
* [Playwright for Go](https://github.com/playwright-community/playwright-go)

The bridge between Node.js and the other languages is basically a Node.js runtime combined with Playwright which gets shipped for each of these languages (around 50MB) and then communicates over stdio to send the relevant commands. This will also download the pre-compiled browsers.

## Is Playwright for Go ready?

We are ready for your feedback, but we are still covering Playwright Go with the tests.

## Resources

* [Playwright for Go Documentation](https://pkg.go.dev/github.com/playwright-community/playwright-go)
* [Playwright Documentation](https://playwright.dev/docs/api/class-playwright)
* [Example recipes](https://github.com/playwright-community/playwright-go/tree/main/examples)
