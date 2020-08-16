package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsConnected(t *testing.T) {
	pw, err := Run()
	require.NoError(t, err)
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	require.True(t, browser.IsConnected)
}

func TestVersion(t *testing.T) {
	pw, err := Run()
	require.NoError(t, err)
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	require.Greater(t, len(browser.Version()), 2)
}

func TestNewContext(t *testing.T) {
	pw, err := Run()
	require.NoError(t, err)
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	context, err := browser.NewContext()
	require.NoError(t, err)
	require.Equal(t, len(context.Pages), 0)
}
