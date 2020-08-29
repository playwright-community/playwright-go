package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/mxschmitt/playwright-go"
)

func exitIfErrorf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func assertEqual(expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		panic(fmt.Sprintf("%v does not equal %v", actual, expected))
	}
}

const TODO_NAME = "Bake a cake"

func main() {
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
	_, err = page.Goto("http://todomvc.com/examples/react/")
	exitIfErrorf("could not goto: %v", err)

	// Helper function to get the amount of todos on the page
	assertCountOfTodos := func(shouldBeCount int) {
		count, err := page.EvaluateOnSelectorAll("ul.todo-list > li", "el => el.length")
		exitIfErrorf("could not determine todo list count: %v", err)
		assertEqual(shouldBeCount, count)
	}

	// Initially there should be 0 entries
	assertCountOfTodos(0)

	// Adding a todo entry (click in the input, enter the todo title and press the Enter key)
	exitIfErrorf("could not click: %v", page.Click("input.new-todo"))
	exitIfErrorf("could not type: %v", page.Type("input.new-todo", TODO_NAME))
	exitIfErrorf("could not press: %v", page.Press("input.new-todo", "Enter"))

	// After adding 1 there should be 1 entry in the list
	assertCountOfTodos(1)

	// Here we get the text in the first todo item to see if it"s the same which the user has entered
	textContentOfFirstTodoEntry, err := page.EvaluateOnSelector("ul.todo-list > li:nth-child(1) label", "el => el.textContent")
	exitIfErrorf("could not get first todo entry: %v", err)
	assertEqual(TODO_NAME, textContentOfFirstTodoEntry)

	// The todo list should be persistent. Here we reload the page and see if the entry is still there
	_, err = page.Reload()
	exitIfErrorf("could not reload: %v", err)
	assertCountOfTodos(1)

	// Set the entry to completed
	exitIfErrorf("could not click: %v", page.Click("input.toggle"))

	// Filter for active entries. There should be 0, because we have completed the entry already
	exitIfErrorf("could not click: %v", page.Click("text=Active"))
	assertCountOfTodos(0)

	// If we filter now for completed entries, there should be 1
	exitIfErrorf("could not click: %v", page.Click("text=Completed"))
	assertCountOfTodos(1)

	// Clear the list of completed entries, then it should be again 0
	exitIfErrorf("could not click: %v", page.Click("text=Clear completed"))
	assertCountOfTodos(0)

	err = browser.Close()
	exitIfErrorf("could not close browser: %v", err)
	err = pw.Stop()
	exitIfErrorf("could not stop Playwright: %v", err)
}
