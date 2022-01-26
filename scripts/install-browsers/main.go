//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func main() {
	if err := playwright.Install(); err != nil {
		log.Fatalf("could not install playwright: %v", err)
	}
}
