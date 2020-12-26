package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDialog(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	helper.Page.On("dialog", func(dialog *dialogImpl) {
		require.Equal(t, "alert", dialog.Type())
		require.Equal(t, "", dialog.DefaultValue())
		require.Equal(t, "yo", dialog.Message())
		require.NoError(t, dialog.Accept())
	})
	_, err := helper.Page.Evaluate("alert('yo')")
	require.NoError(t, err)
}
