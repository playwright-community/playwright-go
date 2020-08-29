package playwright

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadBasic(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	helper.server.SetRoute("/downloadWithFilename", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("Content-Disposition", "attachment; filename=file.txt")
		if _, err := w.Write([]byte("foobar")); err != nil {
			log.Printf("could not write: %v", err)
		}
	})
	require.NoError(t, helper.Page.SetContent(
		fmt.Sprintf(`<a href="%s/downloadWithFilename">download</a>`, helper.server.PREFIX),
	))

	download, err := helper.Page.ExpectDownload(func() error {
		return helper.Page.Click("a")
	})
	require.NoError(t, err)
	require.Equal(t, download.URL(), fmt.Sprintf("%s/downloadWithFilename", helper.server.PREFIX))
	require.Equal(t, download.SuggestedFilename(), "file.txt")
	require.Equal(t, download.String(), "file.txt")
	require.Nil(t, download.Failure())

	file, err := download.Path()
	require.NoError(t, err)
	require.FileExists(t, file)

	tmpFile := filepath.Join(t.TempDir(), download.SuggestedFilename())
	require.NoFileExists(t, tmpFile)
	require.NoError(t, download.SaveAs(tmpFile))
	require.FileExists(t, tmpFile)

	require.NoError(t, download.Delete())
	require.NoFileExists(t, file)
}
