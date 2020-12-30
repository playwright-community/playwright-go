package playwright_test

import (
	"testing"

	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestDialog(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	page.On("dialog", func(dialog playwright.Dialog) {
		require.Equal(t, "alert", dialog.Type())
		require.Equal(t, "", dialog.DefaultValue())
		require.Equal(t, "yo", dialog.Message())
		require.NoError(t, dialog.Accept())
	})
	_, err := page.Evaluate("alert('yo')")
	require.NoError(t, err)
}
