package playwright

import (
	"fmt"
	"log"
	"net/http"
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
	helper.Page.SetContent(
		fmt.Sprintf(`<a href="%s/downloadWithFilename">download</a>`, helper.server.PREFIX),
	)
	// TODO: waitForEvent wrapper
	downloadChan := make(chan *Download, 1)
	helper.Page.On("download", func(ev ...interface{}) {
		downloadChan <- ev[0].(*Download)
	})
	err := helper.Page.Click("a")
	require.NoError(t, err)
	download := <-downloadChan
	require.Equal(t, download.URL(), fmt.Sprintf("%s/downloadWithFilename", helper.server.PREFIX))
	require.Equal(t, download.SuggestedFilename(), "file.txt")
}
