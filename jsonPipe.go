package playwright

import (
	"encoding/json"
	"log"
)

type jsonPipe struct {
	channelOwner
}

func (j *jsonPipe) Send(message map[string]interface{}) error {
	_, err := j.channel.Send("send", map[string]interface{}{
		"message": message})
	return err
}

func (j *jsonPipe) Close() error {
	_, err := j.channel.Send("close")
	return err
}
func newJsonPipe(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *jsonPipe {
	j := &jsonPipe{}
	j.createChannelOwner(j, parent, objectType, guid, initializer)
	j.channel.On("message", func(ev map[string]interface{}) {
		m, err := json.Marshal(ev["message"])
		if err != nil {
			log.Fatal(err)
		}
		var msg message
		err = json.Unmarshal(m, &msg)
		if err != nil {
			log.Fatal(err)
		}
		j.Emit("message", &msg)
	})
	j.channel.On("closed", func() {
		j.Emit("closed")
	})
	return j
}
