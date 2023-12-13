package playwright_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldNotHangWhenPlaywrightUnexpectedExit(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t, false)
	if !isChromium {
		t.Skip("browser agnostic testing")
		return
	}
	if runtime.GOOS == "linux" {
		t.Skip("killPlaywrightProcess not work on linux")
		return
	}

	err := killPlaywrightProcess()
	require.NoError(t, err)
	defer BeforeAll() // need restart playwright driver
	_, err = browser.NewContext()
	require.Error(t, err)
}
