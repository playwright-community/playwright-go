package playwright

import "encoding/base64"

type ElementHandle struct {
	JSHandle
}

func (e *ElementHandle) QuerySelector(selector string) (*ElementHandle, error) {
	channel, err := e.channel.Send("querySelector", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(channel).(*ElementHandle), nil
}

func (e *ElementHandle) TextContent() (string, error) {
	textContent, err := e.channel.Send("textContent")
	if err != nil {
		return "", err
	}
	return textContent.(string), nil
}

func (e *ElementHandle) InnerText() (string, error) {
	innerText, err := e.channel.Send("innerText")
	if err != nil {
		return "", err
	}
	return innerText.(string), nil
}

func (e *ElementHandle) InnerHTML() (string, error) {
	innerHTML, err := e.channel.Send("innerHTML")
	if err != nil {
		return "", err
	}
	return innerHTML.(string), nil
}

func (e *ElementHandle) SetInputFiles(files []InputFile, options ...ElementHandleSetInputFilesOptions) error {
	_, err := e.channel.Send("setInputFiles", map[string]interface{}{
		"files": normalizeFilePayloads(files),
	}, options)
	return err
}

func newElementHandle(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *ElementHandle {
	bt := &ElementHandle{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}

func normalizeFilePayloads(files []InputFile) []map[string]string {
	out := make([]map[string]string, 0)
	for _, file := range files {
		// file.Buffer
		out = append(out, map[string]string{
			"name":     file.Name,
			"mimeType": file.MimeType,
			"buffer":   base64.StdEncoding.EncodeToString(file.Buffer),
		})
	}
	return out
}
