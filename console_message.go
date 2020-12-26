package playwright

type ConsoleMessage struct {
	ChannelOwner
}

type ConsoleMessageLocation struct {
	URL          string `json:"url"`
	LineNumber   int    `json:"lineNumber"`
	ColumnNumber int    `json:"columnNumber"`
}

func (c *ConsoleMessage) Type() string {
	return c.initializer["type"].(string)
}

func (c *ConsoleMessage) Text() string {
	return c.initializer["text"].(string)
}

func (c *ConsoleMessage) String() string {
	return c.Text()
}

func (c *ConsoleMessage) Args() []JSHandleI {
	args := c.initializer["args"].([]interface{})
	out := []JSHandleI{}
	for idx := range args {
		out = append(out, fromChannel(args[idx]).(*JSHandle))
	}
	return out
}

func (c *ConsoleMessage) Location() ConsoleMessageLocation {
	locations := ConsoleMessageLocation{}
	remapMapToStruct(c.initializer["location"], &locations)
	return locations
}

func newConsoleMessage(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *ConsoleMessage {
	bt := &ConsoleMessage{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
