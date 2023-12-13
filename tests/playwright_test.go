package playwright_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldNotHangWhenPlaywrightUnexpectedExit(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t, false)
	defer BeforeAll() // need restart playwright driver
	if !isChromium {
		t.Skip("browser agnostic testing")
		return
	}
	if runtime.GOOS == "linux" {
		t.Skip("ignore linux, hard to find the playwright process")
		return
	}

	err := killPlaywrightProcess()
	require.NoError(t, err)
	_, err = browser.NewContext()
	require.Error(t, err)
}
