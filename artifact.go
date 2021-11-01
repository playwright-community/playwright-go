package playwright

import "errors"

type artifactImpl struct {
	channelOwner
}

func (a *artifactImpl) AbsolutePath() string {
	return a.initializer["absolutePath"].(string)
}

func (a *artifactImpl) PathAfterFinished() (string, error) {
	if a.connection.isRemote {
		return "", errors.New("Path is not available when connecting remotely. Use SaveAs() to save a local copy")
	}
	path, err := a.channel.Send("pathAfterFinished")
	return path.(string), err
}

func (a *artifactImpl) SaveAs(path string) error {
	if !a.connection.isRemote {
		_, err := a.channel.Send("saveAs", map[string]interface{}{
			"path": path,
		})
		return err
	}
	streamChannel, err := a.channel.Send("saveAsStream")
	if err != nil {
		return err
	}
	stream := streamChannel.(*streamImpl)
	return stream.SaveAs(path)
}

func (d *artifactImpl) Failure() (string, error) {
	failure, err := d.channel.Send("failure")
	if failure == nil {
		return "", err
	}
	return failure.(string), err
}

func (d *artifactImpl) Delete() error {
	_, err := d.channel.Send("delete")
	return err
}

func (d *artifactImpl) Cancel() error {
	_, err := d.channel.Send("cancel")
	return err
}

func newArtifact(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *artifactImpl {
	artifact := &artifactImpl{}
	artifact.createChannelOwner(artifact, parent, objectType, guid, initializer)
	return artifact
}
