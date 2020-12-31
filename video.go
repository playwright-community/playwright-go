package playwright

import (
	"path/filepath"
)

type videoImpl struct {
	page     *pageImpl
	path     string
	pathChan chan string
}

func (v *videoImpl) Path() string {
	if v.path == "" {
		v.path = <-v.pathChan
	}
	return v.path
}

func (v *videoImpl) setRelativePath(relativePath string) {
	absolutePath := filepath.Join(*v.page.browserContext.options.RecordVideo.Dir, relativePath)
	v.pathChan <- absolutePath
}

func newVideo(page *pageImpl) *videoImpl {
	return &videoImpl{
		page:     page,
		pathChan: make(chan string, 1),
	}
}
