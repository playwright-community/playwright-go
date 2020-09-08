package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBrowserName(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	require.Equal(t, helper.Playwright.Chromium.Name(), "chromium")
	require.Equal(t, helper.Playwright.Firefox.Name(), "firefox")
	require.Equal(t, helper.Playwright.WebKit.Name(), "webkit")
}

func TestExecutablePath(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	require.Greater(t, len(helper.Playwright.Chromium.ExecutablePath()), 0)
}
