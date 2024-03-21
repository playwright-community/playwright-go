package playwright_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestLocatorScreenshotShouldWork(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, page.SetViewportSize(500, 500))
	_, err := page.Goto(server.PREFIX + "/grid.html")
	require.NoError(t, err)
	_, err = page.Evaluate(`window.scrollBy(50, 100)`)
	require.NoError(t, err)
	screenshot, err := page.Locator(".box:nth-of-type(3)").Screenshot()
	require.NoError(t, err)
	require.NotEmpty(t, screenshot)
	AssertToBeGolden(t, screenshot, "screenshot-element-bounding-box.png")
}

func TestShouldScreenshotWithMask(t *testing.T) {
	BeforeEach(t)

	require.NoError(t, page.SetViewportSize(500, 500))
	_, err := page.Goto(server.PREFIX + "/grid.html")
	require.NoError(t, err)

	screenshot, err := page.Screenshot(playwright.PageScreenshotOptions{
		Mask: []playwright.Locator{
			page.Locator("div").Nth(5),
		},
	})
	require.NoError(t, err)
	AssertToBeGolden(t, screenshot, "mask-should-work.png")

	screenshot, err = page.Locator("body").Screenshot(playwright.LocatorScreenshotOptions{
		Mask: []playwright.Locator{
			page.Locator("div").Nth(5),
		},
	})
	require.NoError(t, err)
	AssertToBeGolden(t, screenshot, "mask-should-work-with-locator.png")

	//nolint:staticcheck
	element, err := page.QuerySelector("body")
	require.NoError(t, err)
	//nolint:staticcheck
	screenshot, err = element.Screenshot(playwright.ElementHandleScreenshotOptions{
		Mask: []playwright.Locator{
			page.Locator("div").Nth(5),
		},
	})
	require.NoError(t, err)
	AssertToBeGolden(t, screenshot, "mask-should-work-with-elementhandle.png")
}
