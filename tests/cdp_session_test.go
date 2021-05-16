package playwright_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCDPSessionSend(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	cdpSession, err := browser.NewBrowserCDPSession()
	require.NoError(t, err)
	result, err := cdpSession.Send("Target.getTargets", nil)
	require.NoError(t, err)
	targetInfos := result.(map[string]interface{})["targetInfos"].([]interface{})
	require.Equal(t, 1, len(targetInfos))
}

func TestCDPSessionOn(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	cdpSession, err := page.Context().NewCDPSession(page)
	require.NoError(t, err)
	_, err = cdpSession.Send("Console.enable", nil)
	require.NoError(t, err)
	cdpSession.On("Console.messageAdded", func(params map[string]interface{}) {
		require.NotNil(t, params)
	})
	page.Evaluate(`console.log("hello")`)
	require.NoError(t, cdpSession.Detach())
}

func TestCDPSessionDetach(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	cdpSession, err := browser.NewBrowserCDPSession()
	require.NoError(t, err)
	require.NoError(t, cdpSession.Detach())
}
