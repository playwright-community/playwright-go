package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsConnected(t *testing.T) {
	helper := NewTestHelper(t)
	require.True(t, helper.Browser.IsConnected)
	helper.Close(t)
}

func TestVersion(t *testing.T) {
	helper := NewTestHelper(t)
	require.Greater(t, len(helper.Browser.Version()), 2)
	helper.Close(t)
}

func TestNewContext(t *testing.T) {
	helper := NewTestHelper(t)
	require.Equal(t, 0, len(helper.Context.Pages))
	helper.Close(t)
}
