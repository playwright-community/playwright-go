package playwright

import "errors"

type Download struct {
	ChannelOwner
}

func (d *Download) String() string {
	return d.SuggestedFilename()
}

func (d *Download) URL() string {
	return d.initializer["url"].(string)
}

func (d *Download) SuggestedFilename() string {
	return d.initializer["suggestedFilename"].(string)
}

func (d *Download) Delete() error {
	_, err := d.channel.Send("delete")
	return err
}

func (d *Download) Failure() error {
	path, err := d.channel.Send("failure")
	if err != nil {
		return err
	}
	if path == nil {
		return nil
	}
	return errors.New(path.(string))
}

func (d *Download) Path() (string, error) {
	path, err := d.channel.Send("path")
	return path.(string), err
}

func (d *Download) SaveAs(path string) error {
	_, err := d.channel.Send("saveAs", map[string]interface{}{
		"path": path,
	})
	return err
}

func newDownload(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Download {
	bt := &Download{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
