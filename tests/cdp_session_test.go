package playwright_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCDPSessionSend(t *testing.T) {
	BeforeEach(t)

	cdpSession, err := browser.NewBrowserCDPSession()
	if isChromium {
		require.NoError(t, err)
		result, err := cdpSession.Send("Target.getTargets", nil)
		require.NoError(t, err)
		targetInfos := result.(map[string]interface{})["targetInfos"].([]interface{})
		require.GreaterOrEqual(t, len(targetInfos), 1)
	} else {
		require.Error(t, err)
	}
}

func TestCDPSessionOn(t *testing.T) {
	BeforeEach(t)

	cdpSession, err := page.Context().NewCDPSession(page)
	if isChromium {
		require.NoError(t, err)
		_, err = cdpSession.Send("Console.enable", nil)
		require.NoError(t, err)
		cdpSession.On("Console.messageAdded", func(params map[string]interface{}) {
			require.NotNil(t, params)
		})
		_, err = page.Evaluate(`console.log("hello")`)
		require.NoError(t, err)
		require.NoError(t, cdpSession.Detach())
	} else {
		require.Error(t, err)
	}
}

func TestCDPSessionDetach(t *testing.T) {
	BeforeEach(t)

	cdpSession, err := browser.NewBrowserCDPSession()
	if isChromium {
		require.NoError(t, err)
		require.NoError(t, cdpSession.Detach())
	} else {
		require.Error(t, err)
	}
}
