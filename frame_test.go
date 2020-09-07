package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFrameWaitForNavigationShouldWork(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	response, err := helper.Page.ExpectNavigation(func() error {
		_, err := helper.Page.Evaluate("url => window.location.href = url", helper.server.PREFIX+"/grid.html")
		return err
	})
	require.NoError(t, err)
	require.True(t, response.Ok())
	require.Contains(t, response.URL(), "grid.html")
}

func TestFrameWaitForNavigationAnchorLinks(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent(`<a href="#foobar">foobar</a>`))
	response, err := helper.Page.ExpectNavigation(func() error {
		return helper.Page.Click("a")
	})
	require.NoError(t, err)
	require.Nil(t, response)
	require.Equal(t, helper.server.EMPTY_PAGE+"#foobar", helper.Page.URL())
}
