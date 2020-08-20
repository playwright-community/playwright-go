package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBrowserName(t *testing.T) {
	helper := NewTestHelper(t)
	require.Equal(t, helper.Playwright.Chromium.Name(), "chromium")
	require.Equal(t, helper.Playwright.Firefox.Name(), "firefox")
	require.Equal(t, helper.Playwright.WebKit.Name(), "webkit")
	helper.Close(t)
}

func TestExecutablePath(t *testing.T) {
	helper := NewTestHelper(t)
	require.Greater(t, len(helper.Playwright.Chromium.ExecutablePath()), 0)
	helper.Close(t)
}
