package playwright

type fileChooserImpl struct {
	page          Page
	elementHandle ElementHandle
	isMultiple    bool
}

func (f *fileChooserImpl) Page() Page {
	return f.page
}

func (f *fileChooserImpl) Element() ElementHandle {
	return f.elementHandle
}

func (f *fileChooserImpl) IsMultiple() bool {
	return f.isMultiple
}

// InputFile represents the input file for:
// - FileChooser.SetFiles()
// - ElementHandle.SetInputFiles()
// - Page.SetInputFiles()
type InputFile struct {
	Name     string
	MimeType string
	Buffer   []byte
}

func (f *fileChooserImpl) SetFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error {
	return f.elementHandle.SetInputFiles(files, options...)
}

func newFileChooser(page Page, elementHandle ElementHandle, isMultiple bool) *fileChooserImpl {
	return &fileChooserImpl{
		page:          page,
		elementHandle: elementHandle,
		isMultiple:    isMultiple,
	}
}
