package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsConnected(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.True(t, helper.Browser.IsConnected)
}

func TestVersion(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.Greater(t, len(helper.Browser.Version()), 2)
}

func TestNewContext(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	require.Equal(t, 0, len(helper.Context.Pages))
}
