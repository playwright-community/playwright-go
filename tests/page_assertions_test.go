package playwright_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPageAssertionsToHaveTitle(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, page.SetContent(`<head><title>TEST</title></head>`))

	require.NoError(t, assertions.ExpectPage(page).ToHaveTitle("TEST"))
	require.NoError(t, assertions.ExpectPage(page).ToHaveTitle(regexp.MustCompile("(?i)test")))
	require.NoError(t, assertions.ExpectPage(page).NotToHaveTitle("NOT TEST"))
}

func TestPageAssertionsToHaveURL(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(server.EMPTY_PAGE)
	require.NoError(t, err)

	require.NoError(t, assertions.ExpectPage(page).ToHaveURL(server.EMPTY_PAGE))
	require.NoError(t, assertions.ExpectPage(page).ToHaveURL(regexp.MustCompile(`.*\.html`)))
	require.NoError(t, assertions.ExpectPage(page).NotToHaveURL("https://playwright.dev"))
}
