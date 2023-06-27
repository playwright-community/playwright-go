package playwright

type consoleMessageImpl struct {
	channelOwner
	page Page
}

// ConsoleMessageLocation represents where a console message was logged in the browser
type ConsoleMessageLocation struct {
	URL          string `json:"url"`
	LineNumber   int    `json:"lineNumber"`
	ColumnNumber int    `json:"columnNumber"`
}

func (c *consoleMessageImpl) Type() string {
	return c.initializer["type"].(string)
}

func (c *consoleMessageImpl) Text() string {
	return c.initializer["text"].(string)
}

func (c *consoleMessageImpl) String() string {
	return c.Text()
}

func (c *consoleMessageImpl) Args() []JSHandle {
	args := c.initializer["args"].([]interface{})
	out := []JSHandle{}
	for idx := range args {
		out = append(out, fromChannel(args[idx]).(*jsHandleImpl))
	}
	return out
}

func (c *consoleMessageImpl) Location() ConsoleMessageLocation {
	locations := ConsoleMessageLocation{}
	remapMapToStruct(c.initializer["location"], &locations)
	return locations
}

func (c *consoleMessageImpl) Page() Page {
	return c.page
}

func newConsoleMessage(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *consoleMessageImpl {
	bt := &consoleMessageImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	// Note: currently, we only report console messages for pages and they always have a page.
	// However, in the future we might report console messages for service workers or something else,
	// where page() would be null.
	page := fromNullableChannel(initializer["page"])
	if page != nil {
		bt.page = page.(*pageImpl)
	}
	return bt
}
