package playwright

import (
	"errors"
	"sync"
)

type videoImpl struct {
	page         *pageImpl
	artifact     *artifactImpl
	artifactChan chan *artifactImpl
	closeOnce    sync.Once
	isRemote     bool
}

func (v *videoImpl) Path() (string, error) {
	if v.isRemote {
		return "", errors.New("Path is not available when connecting remotely. Use SaveAs() to save a local copy.")
	}
	v.getArtifact()
	if v.artifact == nil {
		return "", errors.New("Page did not produce any video frames")
	}
	return v.artifact.AbsolutePath(), nil
}

func (v *videoImpl) Delete() error {
	v.getArtifact()
	if v.artifact == nil {
		return nil
	}
	return v.artifact.Delete()
}

func (v *videoImpl) SaveAs(path string) error {
	if !v.page.IsClosed() {
		return errors.New("Page is not yet closed. Close the page prior to calling SaveAs")
	}
	v.getArtifact()
	if v.artifact == nil {
		return errors.New("Page did not produce any video frames")
	}
	return v.artifact.SaveAs(path)
}

func (v *videoImpl) artifactReady(artifact *artifactImpl) {
	v.artifactChan <- artifact
}

func (v *videoImpl) pageClosed() {
	v.closeOnce.Do(func() {
		if v.artifactChan != nil {
			close(v.artifactChan)
		}
	})
}

func (v *videoImpl) getArtifact() {
	// prevent channel block if no video will be produced
	if v.page.browserContext.options == nil {
		v.pageClosed()
	} else {
		option := v.page.browserContext.options
		if option == nil || option.RecordVideo == nil || option.RecordVideo.Dir == nil { // no recordVideo option
			v.pageClosed()
		}
	}
	artifact := <-v.artifactChan
	if artifact != nil {
		v.artifact = artifact
	}
}

func newVideo(page *pageImpl) *videoImpl {
	video := &videoImpl{
		page:     page,
		isRemote: page.connection.isRemote,
	}
	video.artifactChan = make(chan *artifactImpl, 1)
	if page.IsClosed() {
		video.pageClosed()
	} else {
		page.On("close", video.pageClosed)
	}
	return video
}
