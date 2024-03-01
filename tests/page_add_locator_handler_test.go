package playwright_test

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestPageAddLocatorHandlerShouldWork(t *testing.T) {
	BeforeEach(t, playwright.BrowserNewContextOptions{
		HasTouch: playwright.Bool(false), // (v1.42.0) firefox 123.0 doesn't respond to pointerover when hasTouch is true
	})
	defer AfterEach(t)
	_, err := page.Goto(fmt.Sprintf("%s/input/handle-locator.html", server.PREFIX))
	require.NoError(t, err)

	beforeCount := 0
	afterCount := 0

	err = page.AddLocatorHandler(page.GetByText("This interstitial covers the button"), func() {
		beforeCount++
		require.NoError(t, page.Locator("#close").Click())
		afterCount++
	})
	require.NoError(t, err)

	for _, args := range [][]interface{}{
		{"mouseover", 1},
		{"mouseover", 1, "capture"},
		{"mouseover", 2},
		{"mouseover", 2, "capture"},
		{"pointerover", 1},
		{"pointerover", 1, "capture"},
		{"none", 1},
		{"remove", 1},
		{"hide", 1},
	} {
		err = page.Locator("#aside").Hover()
		require.NoError(t, err)
		beforeCount = 0
		afterCount = 0
		_, err = page.Evaluate(`(args) => { window.clicked = 0; window.setupAnnoyingInterstitial(...args); }`, args)
		require.NoError(t, err)
		require.Equal(t, 0, beforeCount)
		require.Equal(t, 0, afterCount)

		err = page.Locator("#target").Click()
		require.NoError(t, err)
		require.Equal(t, args[1].(int), beforeCount)
		require.Equal(t, args[1].(int), afterCount)

		ret, err := page.Evaluate(`window.clicked`)
		require.NoError(t, err)
		require.Equal(t, 1, ret)
		require.NoError(t, expect.Locator(page.Locator(`#interstitial`)).Not().ToBeVisible())
	}
}

func TestPageAddLocatorHandlerShouldWorkWithACustomCheck(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(fmt.Sprintf("%s/input/handle-locator.html", server.PREFIX))
	require.NoError(t, err)

	err = page.AddLocatorHandler(page.Locator("body"), func() {
		ret, _ := page.GetByText("This interstitial covers the button").IsVisible()
		if ret {
			require.NoError(t, page.Locator("#close").Click())
		}
	})
	require.NoError(t, err)

	for _, args := range [][]interface{}{
		{"mouseover", 2},
		{"none", 1},
		{"remove", 1},
		{"hide", 1},
	} {
		err = page.Locator("#aside").Hover()
		require.NoError(t, err)

		_, err = page.Evaluate(`(args) => { window.clicked = 0; window.setupAnnoyingInterstitial(...args); }`, args)
		require.NoError(t, err)
		err = page.Locator("#target").Click()
		require.NoError(t, err)
		ret, err := page.Evaluate(`window.clicked`)
		require.NoError(t, err)
		require.Equal(t, 1, ret)
		require.NoError(t, expect.Locator(page.Locator(`#interstitial`)).Not().ToBeVisible())
	}
}

func TestPageAddLocatorHandlerShouldWorkWithLocatorHover(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(fmt.Sprintf("%s/input/handle-locator.html", server.PREFIX))
	require.NoError(t, err)

	err = page.AddLocatorHandler(page.GetByText("This interstitial covers the button"), func() {
		require.NoError(t, page.Locator("#close").Click())
	})
	require.NoError(t, err)

	err = page.Locator("#aside").Hover()
	require.NoError(t, err)
	_, err = page.Evaluate(`() => { window.setupAnnoyingInterstitial("pointerover", 1, "capture"); }`, nil)
	require.NoError(t, err)
	err = page.Locator("#target").Hover()
	require.NoError(t, err)
	require.NoError(t, expect.Locator(page.Locator(`#interstitial`)).Not().ToBeVisible())
	// nolint:staticcheck
	ret, err := page.EvalOnSelector("#target", `e => window.getComputedStyle(e).backgroundColor`, nil)
	require.NoError(t, err)
	require.Equal(t, "rgb(255, 255, 0)", ret)
}

func TestPageAddLocatorHandlerShouldWorkWithForceTrue(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(fmt.Sprintf("%s/input/handle-locator.html", server.PREFIX))
	require.NoError(t, err)

	err = page.AddLocatorHandler(page.GetByText("This interstitial covers the button"), func() {
		require.NoError(t, page.Locator("#close").Click())
	})
	require.NoError(t, err)

	err = page.Locator("#aside").Hover()
	require.NoError(t, err)
	_, err = page.Evaluate(`() => { window.setupAnnoyingInterstitial("none", 1); }`, nil)
	require.NoError(t, err)
	err = page.Locator("#target").Click(playwright.LocatorClickOptions{
		Force:   playwright.Bool(true),
		Timeout: playwright.Float(2000),
	})
	require.NoError(t, err)
	visible, err := page.Locator(`#interstitial`).IsVisible()
	require.NoError(t, err)
	require.True(t, visible)
	ret, err := page.Evaluate(`window.clicked`)
	require.NoError(t, err)
	require.Equal(t, nil, ret)
}

func TestPageAddLocatorHandlerShouldThrowWhenPageCloses(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(fmt.Sprintf("%s/input/handle-locator.html", server.PREFIX))
	require.NoError(t, err)

	err = page.AddLocatorHandler(page.GetByText("This interstitial covers the button"), func() {
		require.NoError(t, page.Close())
	})
	require.NoError(t, err)

	err = page.Locator("#aside").Hover()
	require.NoError(t, err)
	_, err = page.Evaluate(`() => { window.clicked = 0; window.setupAnnoyingInterstitial("mouseover", 1); }`, nil)
	require.NoError(t, err)
	err = page.Locator("#target").Click()
	require.ErrorIs(t, err, playwright.ErrTargetClosed)
}

func TestPageAddLocatorHandlerShouldThrowWhenHandlerTimesOut(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(fmt.Sprintf("%s/input/handle-locator.html", server.PREFIX))
	require.NoError(t, err)

	called := atomic.Int32{}
	stallChan := make(chan struct{})

	err = page.AddLocatorHandler(page.GetByText("This interstitial covers the button"), func() {
		called.Add(1)
		// Deliberately timeout.
		<-stallChan
	})
	require.NoError(t, err)

	err = page.Locator("#aside").Hover()
	require.NoError(t, err)
	_, err = page.Evaluate(`() => { window.clicked = 0; window.setupAnnoyingInterstitial("mouseover", 1); }`, nil)
	require.NoError(t, err)
	err = page.Locator("#target").Click(playwright.LocatorClickOptions{
		Timeout: playwright.Float(2000),
	})
	require.ErrorIs(t, err, playwright.ErrTimeout)
	err = page.Locator("#target").Click(playwright.LocatorClickOptions{
		Timeout: playwright.Float(2000),
	})
	require.ErrorIs(t, err, playwright.ErrTimeout)
	require.Equal(t, int32(1), called.Load())
	stallChan <- struct{}{}
}

func TestPageAddLocatorHandlerShouldWorkWithToBeVisible(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	_, err := page.Goto(fmt.Sprintf("%s/input/handle-locator.html", server.PREFIX))
	require.NoError(t, err)

	called := 0

	err = page.AddLocatorHandler(page.GetByText("This interstitial covers the button"), func() {
		called++
		require.NoError(t, page.Locator("#close").Click())
	})
	require.NoError(t, err)

	_, err = page.Evaluate(`() => { window.clicked = 0; window.setupAnnoyingInterstitial("remove", 1); }`, nil)
	require.NoError(t, err)
	require.NoError(t, expect.Locator(page.Locator(`#target`)).ToBeVisible())
	require.NoError(t, expect.Locator(page.Locator(`#interstitial`)).Not().ToBeVisible())
	require.Equal(t, 1, called)
}
