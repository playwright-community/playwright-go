//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func assertEqual(expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		panic(fmt.Sprintf("%v does not equal %v", actual, expected))
	}
}

const todoName = "Bake a cake"

func main() {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)
	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)
	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)
	_, err = page.Goto("http://todomvc.com/examples/react/")
	assertErrorToNilf("could not goto: %w", err)

	// Helper function to get the amount of todos on the page
	assertCountOfTodos := func(shouldBeCount int) {
		count, err := page.EvalOnSelectorAll("ul.todo-list > li", "el => el.length")
		assertErrorToNilf("could not determine todo list count: %w", err)
		assertEqual(shouldBeCount, count)
	}

	// Initially there should be 0 entries
	assertCountOfTodos(0)

	// Adding a todo entry (click in the input, enter the todo title and press the Enter key)
	assertErrorToNilf("could not click: %v", page.Click("input.new-todo"))
	assertErrorToNilf("could not type: %v", page.Type("input.new-todo", todoName))
	assertErrorToNilf("could not press: %v", page.Press("input.new-todo", "Enter"))

	// After adding 1 there should be 1 entry in the list
	assertCountOfTodos(1)

	// Here we get the text in the first todo item to see if it"s the same which the user has entered
	textContentOfFirstTodoEntry, err := page.EvalOnSelector("ul.todo-list > li:nth-child(1) label", "el => el.textContent")
	assertErrorToNilf("could not get first todo entry: %w", err)
	assertEqual(todoName, textContentOfFirstTodoEntry)

	// The todo list should be persistent. Here we reload the page and see if the entry is still there
	_, err = page.Reload()
	assertErrorToNilf("could not reload: %w", err)
	assertCountOfTodos(1)

	// Set the entry to completed
	assertErrorToNilf("could not click: %v", page.Click("input.toggle"))

	// Filter for active entries. There should be 0, because we have completed the entry already
	assertErrorToNilf("could not click: %v", page.Click("text=Active"))
	assertCountOfTodos(0)

	// If we filter now for completed entries, there should be 1
	assertErrorToNilf("could not click: %v", page.Click("text=Completed"))
	assertCountOfTodos(1)

	// Clear the list of completed entries, then it should be again 0
	assertErrorToNilf("could not click: %v", page.Click("text=Clear completed"))
	assertCountOfTodos(0)

	assertErrorToNilf("could not close browser: %w", browser.Close())
	assertErrorToNilf("could not stop Playwright: %w", pw.Stop())
}
