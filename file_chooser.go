package playwright

type FileChooser struct {
	page          *Page
	elementHandle *ElementHandle
	isMultiple    bool
}

func (f *FileChooser) Page() *Page {
	return f.page
}

func (f *FileChooser) Element() *ElementHandle {
	return f.elementHandle
}

func (f *FileChooser) IsMultiple() bool {
	return f.isMultiple
}

type InputFile struct {
	Name     string
	MimeType string
	Buffer   []byte
}

func (e *FileChooser) SetFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error {
	return e.elementHandle.SetInputFiles(files, options...)
}

func newFileChooser(page *Page, elementHandle *ElementHandle, isMultiple bool) *FileChooser {
	return &FileChooser{
		page:          page,
		elementHandle: elementHandle,
		isMultiple:    isMultiple,
	}
}
