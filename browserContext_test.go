package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPage(t *testing.T) {
	pw, err := Run()
	if err != nil {
		t.Fatalf("could not launch playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch()
	require.NoError(t, err)
	context, err := browser.NewContext()
	require.NoError(t, err)
	page, err := context.NewPage()
	require.NoError(t, err)
	require.NotNil(t, page)
}