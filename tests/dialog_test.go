package playwright_test

import (
	"testing"

	"github.com/playwright-community/playwright-go"
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

func TestDialogDismiss(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	page.On("dialog", func(dialog playwright.Dialog) {
		require.NoError(t, dialog.Dismiss())
	})
	result, err := page.Evaluate("prompt('question?')")
	require.NoError(t, err)
	require.Equal(t, result, nil)
}

func TestDialogAcceptWithText(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	page.On("dialog", func(dialog playwright.Dialog) {
		require.NoError(t, dialog.Accept("hey foobar"))
	})
	result, err := page.Evaluate("prompt('question?')")
	require.NoError(t, err)
	require.Equal(t, result, "hey foobar")
}
