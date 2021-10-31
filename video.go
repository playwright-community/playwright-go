package playwright

import "errors"

type videoImpl struct {
	page     *pageImpl
	artifact *artifactImpl
	isRemote bool
}

func (v *videoImpl) Path() (string, error) {
	if v.isRemote {
		return "", errors.New("Path is not available when connecting remotely. Use SaveAs() to save a local copy.")
	}
	if v.artifact == nil {
		return "", errors.New("Page did not produce any video frames")
	}
	return v.artifact.AbsolutePath(), nil
}

func (v *videoImpl) Delete() error {
	return v.artifact.Delete()
}

func (v *videoImpl) SaveAs(path string) error {
	if v.artifact == nil {
		return errors.New("Page did not produce any video frames")
	}
	return v.artifact.SaveAs(path)
}

func newVideo(page *pageImpl) *videoImpl {
	return &videoImpl{
		page:     page,
		isRemote: page.connection.isRemote,
	}
}
