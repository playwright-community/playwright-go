package playwright

type consoleMessageImpl struct {
	event map[string]interface{}
	page  Page
}

func (c *consoleMessageImpl) Type() string {
	return c.event["type"].(string)
}

func (c *consoleMessageImpl) Text() string {
	return c.event["text"].(string)
}

func (c *consoleMessageImpl) String() string {
	return c.Text()
}

func (c *consoleMessageImpl) Args() []JSHandle {
	args := c.event["args"].([]interface{})
	out := []JSHandle{}
	for idx := range args {
		out = append(out, fromChannel(args[idx]).(*jsHandleImpl))
	}
	return out
}

func (c *consoleMessageImpl) Location() *ConsoleMessageLocation {
	location := &ConsoleMessageLocation{}
	remapMapToStruct(c.event["location"], location)
	return location
}

func (c *consoleMessageImpl) Page() Page {
	return c.page
}

func newConsoleMessage(event map[string]interface{}) *consoleMessageImpl {
	bt := &consoleMessageImpl{}
	bt.event = event
	page := fromNullableChannel(event["page"])
	if page != nil {
		bt.page = page.(*pageImpl)
	}
	return bt
}
