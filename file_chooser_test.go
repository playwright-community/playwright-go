package playwright

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileChooser(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.PREFIX + "/input/fileupload.html")
	require.NoError(t, err)
	input, err := helper.Page.QuerySelector("input")
	require.NoError(t, err)
	file, err := ioutil.ReadFile(helper.Asset("file-to-upload.txt"))
	require.NoError(t, err)
	err = input.SetInputFiles([]InputFile{
		{
			Name:     "file-to-upload.txt",
			MimeType: "text/plain",
			Buffer:   file,
		},
	})
	require.NoError(t, err)
	fileName, err := helper.Page.Evaluate("e => e.files[0].name", input)
	require.NoError(t, err)
	require.Equal(t, "file-to-upload.txt", fileName)
	content, err := helper.Page.Evaluate(`e => {
        reader = new FileReader()
        promise = new Promise(fulfill => reader.onload = fulfill)
        reader.readAsText(e.files[0])
        return promise.then(() => reader.result)
    }`, input)
	require.NoError(t, err)
	require.Equal(t, "contents of the file", content)
}
