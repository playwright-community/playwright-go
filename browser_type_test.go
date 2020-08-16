package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBrowserName(t *testing.T) {
	pw, err := Run()
	require.NoError(t, err)
	require.Equal(t, pw.Chromium.Name(), "chromium")
	require.Equal(t, pw.Firefox.Name(), "firefox")
	require.Equal(t, pw.WebKit.Name(), "webkit")
}

func TestExecutablePath(t *testing.T) {
	pw, err := Run()
	require.NoError(t, err)
	require.Greater(t, len(pw.Chromium.ExecutablePath()), 0)
}
