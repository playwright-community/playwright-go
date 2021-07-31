package playwright

type videoImpl struct {
	page     *pageImpl
	artifact *artifactImpl
}

func (v *videoImpl) Path() string {
	return v.artifact.AbsolutePath()
}

func (v *videoImpl) Delete() error {
	return v.artifact.Delete()
}

func (v *videoImpl) SaveAs(path string) error {
	return v.artifact.SaveAs(path)
}
func (v *videoImpl) setArtifact(artifact *artifactImpl) {
	v.artifact = artifact
}

func newVideo(page *pageImpl) *videoImpl {
	return &videoImpl{
		page: page,
	}
}
