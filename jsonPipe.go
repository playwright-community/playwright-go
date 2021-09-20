package playwright

import (
	"encoding/json"
)

type jsonPipe struct {
	channelOwner
}

func (j *jsonPipe) Send(message map[string]interface{}) error {
	if _, err := j.channel.Send("send", map[string]interface{}{
		"message": message,
	}); err != nil {
		return err
	}
	return nil
}

func (j *jsonPipe) Close() error {
	if _, err := j.channel.Send("close"); err != nil {
		return err
	}
	return nil
}
func newJsonPipe(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *jsonPipe {
	j := &jsonPipe{}
	j.createChannelOwner(j, parent, objectType, guid, initializer)
	j.channel.On("message", func(ev map[string]interface{}) {
		m, _ := json.Marshal(ev["message"])
		var msg message
		_ = json.Unmarshal(m, &msg)
		j.Emit("message", &msg)
	})
	j.channel.On("closed", func() {
		j.Emit("closed")
	})
	return j
}
