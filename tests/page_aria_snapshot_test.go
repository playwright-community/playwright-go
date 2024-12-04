package playwright_test

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func Unshift(snapshot string) string {
	lines := strings.Split(snapshot, "\n")
	whitespacePrefixLength := 100
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		// replace each tab with 2 spaces
		match := regexp.MustCompile(`^(\t*)`).FindStringSubmatch(line)
		if len(match) > 1 {
			lines[i] = regexp.MustCompile(`^(\t*)`).ReplaceAllString(line, strings.Repeat("  ", len(match[1])))
			if len(match[1]) < whitespacePrefixLength {
				whitespacePrefixLength = len(match[1]) * 2
			}
		}
	}

	for i, line := range lines {
		if line == "" {
			continue
		}
		lines[i] = regexp.MustCompile(fmt.Sprintf(`^(\s{0,%d})`, whitespacePrefixLength)).ReplaceAllString(line, "")
	}
	return strings.Join(slices.DeleteFunc(lines, func(line string) bool { return strings.TrimSpace(line) == "" }), "\n")
}

func checkAndMatchSnapshot(t *testing.T, locator playwright.Locator, snapshot string) {
	t.Helper()
	snapshot = Unshift(snapshot)
	ariaSnapshot, err := locator.AriaSnapshot()
	require.NoError(t, err)
	require.Equal(t, snapshot, ariaSnapshot)
	require.NoError(t, expect.Locator(locator).ToMatchAriaSnapshot(snapshot))
}

func TestShouldSnapshot(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, page.SetContent(`<h1>title</h1>`))
	checkAndMatchSnapshot(t, page.Locator("body"), `
	- heading "title" [level=1]
	`)
}

func TestShouldSnapshotList(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, page.SetContent(`<h1>title</h1><h1>title 2</h1>`))
	checkAndMatchSnapshot(t, page.Locator("body"), `
	- heading "title" [level=1]
	- heading "title 2" [level=1]
	`)
}

func TestShouldSnapshotListWithList(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, page.SetContent(`<ul><li>one</li><li>two</li></ul>`))
	checkAndMatchSnapshot(t, page.Locator("body"), `
	- list:
		- listitem: one
		- listitem: two
	`)
}

func TestShouldSnapshotListWithAccessibleName(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, page.SetContent(`<ul aria-label="my list"><li>one</li><li>two</li></ul>`))
	checkAndMatchSnapshot(t, page.Locator("body"), `
	- list "my list":
		- listitem: one
		- listitem: two
	`)
}

func TestShouldSnapshotComplex(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, page.SetContent(`<ul><li><a href="about:blank">link</a></li></ul>`))
	checkAndMatchSnapshot(t, page.Locator("body"), `
	- list:
		- listitem:
			- link "link"
	`)
}
