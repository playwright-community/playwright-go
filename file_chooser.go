package playwright

type filechooserImpl struct {
	page          Page
	elementHandle ElementHandle
	isMultiple    bool
}

func (f *filechooserImpl) Page() Page {
	return f.page
}

func (f *filechooserImpl) Element() ElementHandle {
	return f.elementHandle
}

func (f *filechooserImpl) IsMultiple() bool {
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

func (f *filechooserImpl) SetFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error {
	return f.elementHandle.SetInputFiles(files, options...)
}

func newFileChooser(page Page, elementHandle ElementHandle, isMultiple bool) *filechooserImpl {
	return &filechooserImpl{
		page:          page,
		elementHandle: elementHandle,
		isMultiple:    isMultiple,
	}
}
