package playwright

type FileChooser struct {
	page          PageI
	elementHandle ElementHandleI
	isMultiple    bool
}

func (f *FileChooser) Page() PageI {
	return f.page
}

func (f *FileChooser) Element() ElementHandleI {
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

func newFileChooser(page PageI, elementHandle ElementHandleI, isMultiple bool) *FileChooser {
	return &FileChooser{
		page:          page,
		elementHandle: elementHandle,
		isMultiple:    isMultiple,
	}
}
