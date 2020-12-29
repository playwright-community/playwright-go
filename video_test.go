package playwright_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/h2non/filetype"
	"github.com/mxschmitt/playwright-go"
	"github.com/stretchr/testify/require"
)

func TestVideoShouldWork(t *testing.T) {
	recordVideoDir := t.TempDir()
	helper := newContextWithOptions(t, playwright.BrowserNewContextOptions{
		RecordVideo: &playwright.BrowserNewContextRecordVideo{
			Dir: playwright.String(recordVideoDir),
		},
	})
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	_, err = helper.Page.Reload()
	require.NoError(t, err)
	_, err = helper.Page.Reload()
	require.NoError(t, err)
	require.NoError(t, helper.Context.Close())

	files, err := ioutil.ReadDir(recordVideoDir)
	require.NoError(t, err)
	require.Equal(t, len(files), 1)
	content, err := ioutil.ReadFile(filepath.Join(recordVideoDir, files[0].Name()))
	require.NoError(t, err)
	require.True(t, filetype.IsVideo(content))
}
