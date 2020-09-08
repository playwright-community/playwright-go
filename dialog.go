package playwright

type Dialog struct {
	ChannelOwner
}

func (d *Dialog) Type() string {
	return d.initializer["type"].(string)
}

func (d *Dialog) Message() string {
	return d.initializer["message"].(string)
}

func (d *Dialog) DefaultValue() string {
	return d.initializer["defaultValue"].(string)
}

func (d *Dialog) Accept(texts ...string) error {
	var promptText *string
	if len(texts) == 1 {
		promptText = &texts[0]
	}
	_, err := d.channel.Send("accept", map[string]interface{}{
		"promptText": promptText,
	})
	return err
}

func (d *Dialog) Dismiss() error {
	_, err := d.channel.Send("dismiss")
	return err
}

func newDialog(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Dialog {
	bt := &Dialog{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
