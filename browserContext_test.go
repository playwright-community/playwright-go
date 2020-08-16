package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPage(t *testing.T) {
	pw, err := Run()
	require.NoError(t, err)
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	context, err := browser.NewContext()
	require.NoError(t, err)
	page, err := context.NewPage()
	require.NoError(t, err)
	require.NotNil(t, page)
}
