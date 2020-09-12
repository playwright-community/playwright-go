package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBrowserIsConnected(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	require.True(t, helper.Browser.IsConnected)
}

func TestBrowserVersion(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	require.Greater(t, len(helper.Browser.Version()), 2)
}

func TestBrowserNewContext(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	require.Equal(t, 1, len(helper.Context.Pages()))
}

func TestBrowserNewPage(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	require.Equal(t, 1, len(helper.Browser.Contexts()))
	page, err := helper.Browser.NewPage()
	require.NoError(t, err)
	require.Equal(t, 2, len(helper.Browser.Contexts()))
	require.NoError(t, page.Close())
	require.Equal(t, 1, len(helper.Browser.Contexts()))
}

func TestBrowserClose(t *testing.T) {
	pw, err := Run()
	require.NoError(t, err)
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	require.NoError(t, browser.Close())
	require.NoError(t, pw.Stop())
}
