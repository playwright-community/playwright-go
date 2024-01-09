package playwright

import (
	"errors"
	"sync"
)

type videoImpl struct {
	page         *pageImpl
	artifact     *artifactImpl
	artifactChan chan *artifactImpl
	done         chan struct{}
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

func (v *videoImpl) pageClosed(p Page) {
	v.closeOnce.Do(func() {
		close(v.done)
	})
}

func (v *videoImpl) getArtifact() {
	// prevent channel block if no video will be produced
	if v.page.browserContext.options == nil {
		v.pageClosed(v.page)
	} else {
		option := v.page.browserContext.options
		if option == nil || option.RecordVideo == nil { // no recordVideo option
			v.pageClosed(v.page)
		}
	}
	select {
	case artifact := <-v.artifactChan:
		if artifact != nil {
			v.artifact = artifact
		}
	case <-v.done: // page closed
		select { // make sure get artifact if it's ready before page closed
		case artifact := <-v.artifactChan:
			if artifact != nil {
				v.artifact = artifact
			}
		default:
		}
	}
}

func newVideo(page *pageImpl) *videoImpl {
	video := &videoImpl{
		page:         page,
		artifactChan: make(chan *artifactImpl, 1),
		done:         make(chan struct{}, 1),
		isRemote:     page.connection.isRemote,
	}

	if page.isClosed {
		video.pageClosed(page)
	} else {
		page.OnClose(video.pageClosed)
	}
	return video
}
