package playwright

import "errors"

type downloadImpl struct {
	channelOwner
}

func (d *downloadImpl) String() string {
	return d.SuggestedFilename()
}

func (d *downloadImpl) URL() string {
	return d.initializer["url"].(string)
}

func (d *downloadImpl) SuggestedFilename() string {
	return d.initializer["suggestedFilename"].(string)
}

func (d *downloadImpl) Delete() error {
	_, err := d.channel.Send("delete")
	return err
}

func (d *downloadImpl) Failure() error {
	path, err := d.channel.Send("failure")
	if err != nil {
		return err
	}
	if path == nil {
		return nil
	}
	return errors.New(path.(string))
}

func (d *downloadImpl) Path() (string, error) {
	path, err := d.channel.Send("path")
	return path.(string), err
}

func (d *downloadImpl) SaveAs(path string) error {
	_, err := d.channel.Send("saveAs", map[string]interface{}{
		"path": path,
	})
	return err
}

func newDownload(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *downloadImpl {
	bt := &downloadImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
