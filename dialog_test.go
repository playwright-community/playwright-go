package playwright

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDialog(t *testing.T) {
	helper := BeforeEach(t)
	defer helper.AfterEach()
	helper.Page.On("dialog", func(ev ...interface{}) {
		dialog := ev[0].(*Dialog)
		require.Equal(t, "alert", dialog.Type())
		require.Equal(t, "", dialog.DefaultValue())
		require.Equal(t, "yo", dialog.Message())
		require.NoError(t, dialog.Accept())
		fmt.Println("HIT")
	})
	_, err := helper.Page.Evaluate("alert('yo')")
	require.NoError(t, err)
}
