package playwright_test

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDownloadBasic(t *testing.T) {
	BeforeEach(t)
	defer AfterEach(t)
	server.SetRoute("/downloadWithFilename", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("Content-Disposition", "attachment; filename=file.txt")
		if _, err := w.Write([]byte("foobar")); err != nil {
			log.Printf("could not write: %v", err)
		}
	})
	server.SetRoute("/downloadWithDelay", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/octet-stream")
		w.Header().Add("Content-Disposition", "attachment")
		if _, err := w.Write([]byte(strings.Repeat("foobar", 8192))); err != nil {
			log.Printf("could not write: %v", err)
		}
	})
	require.NoError(t, page.SetContent(
		fmt.Sprintf(`<a href="%s/downloadWithFilename">download</a>`, server.PREFIX),
	))

	download, err := page.ExpectDownload(func() error {
		return page.Click("a")
	})
	require.NoError(t, err)
	require.Equal(t, download.URL(), fmt.Sprintf("%s/downloadWithFilename", server.PREFIX))
	require.Equal(t, download.SuggestedFilename(), "file.txt")
	require.Equal(t, download.String(), "file.txt")
	failure, err := download.Failure()
	require.NoError(t, err)
	require.Equal(t, failure, "")

	file, err := download.Path()
	require.NoError(t, err)
	require.FileExists(t, file)

	tmpFile := filepath.Join(t.TempDir(), download.SuggestedFilename())
	require.NoFileExists(t, tmpFile)
	require.NoError(t, download.SaveAs(tmpFile))
	require.FileExists(t, tmpFile)

	require.NoError(t, download.Delete())
	require.NoFileExists(t, file)

	require.NoError(t, page.SetContent(
		fmt.Sprintf(`<a href="%s/downloadWithDelay">download</a>`, server.PREFIX),
	))
	download, err = page.ExpectDownload(func() error {
		return page.Click("a")
	})
	require.NoError(t, err)
	require.Equal(t, download.URL(), fmt.Sprintf("%s/downloadWithDelay", server.PREFIX))
	err = download.Cancel()
	require.NoError(t, err)
	failure, err = download.Failure()
	require.NoError(t, err)
	require.Equal(t, failure, "canceled")
}
