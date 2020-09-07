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

func TestFileChooserShouldEmitEvent(t *testing.T) {
	helper := NewTestHelper(t)
	defer helper.AfterEach()
	_, err := helper.Page.Goto(helper.server.EMPTY_PAGE)
	require.NoError(t, err)
	require.NoError(t, helper.Page.SetContent("<input type=file>"))

	fileChooser, err := helper.Page.ExpectFileChooser(func() error {
		return helper.Page.Click("input")
	})
	require.NoError(t, err)
	require.False(t, fileChooser.IsMultiple())
	require.Equal(t, helper.Page, fileChooser.Page())
	elementHTML, err := fileChooser.Element().InnerHTML()
	require.NoError(t, err)

	inputElement, err := helper.Page.QuerySelector("input")
	require.NoError(t, err)
	inputElementHTML, err := inputElement.InnerHTML()
	require.NoError(t, err)

	require.Equal(t, elementHTML, inputElementHTML)

	err = fileChooser.setFiles([]InputFile{
		{
			Name:     "file-to-upload.txt",
			MimeType: "text/plain",
			Buffer:   []byte("123"),
		},
	})
	require.NoError(t, err)
	fileName, err := helper.Page.Evaluate("e => e.files[0].name", inputElement)
	require.NoError(t, err)
	require.Equal(t, "file-to-upload.txt", fileName)
}
