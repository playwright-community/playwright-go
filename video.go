package playwright

type videoImpl struct {
	page         *pageImpl
	path         string
	artifactChan chan *artifactImpl
}

func (v *videoImpl) Path() string {
	if v.path == "" {
		v.path = (<-v.artifactChan).AbsolutePath()
	}
	return v.path
}

func (v *videoImpl) setArtifact(artifact *artifactImpl) {
	v.artifactChan <- artifact
}

func newVideo(page *pageImpl) *videoImpl {
	return &videoImpl{
		page:         page,
		artifactChan: make(chan *artifactImpl, 1),
	}
}
