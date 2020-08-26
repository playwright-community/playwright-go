package playwright

import "sync"

type Frame struct {
	ChannelOwner
	sync.RWMutex
	page *Page
	url  string
}

func (b *Frame) URL() string {
	b.RLock()
	defer b.RUnlock()
	return b.url
}

func (b *Frame) SetContent(content string, options ...PageSetContentOptions) error {
	_, err := b.channel.Send("setContent", map[string]interface{}{
		"html": content,
	}, options)
	return err
}

func (b *Frame) Content() (string, error) {
	content, err := b.channel.Send("content")
	return content.(string), err
}

func (b *Frame) Goto(url string) error {
	_, err := b.channel.Send("goto", map[string]interface{}{
		"url": url,
	})
	return err
}

func (b *Frame) Page() *Page {
	return b.page
}

func (b *Frame) onFrameNavigated(event ...interface{}) {
	b.Lock()
	b.url = event[0].(map[string]interface{})["url"].(string)
	b.Unlock()
}

func (b *Frame) QuerySelector(selector string) (*ElementHandle, error) {
	channelOwner, err := b.channel.Send("querySelector", map[string]interface{}{
		"selector": selector,
	})
	if err != nil {
		return nil, err
	}
	return channelOwner.(*Channel).object.(*ElementHandle), nil
}

func (b *Frame) Evaluate(expression string, options ...interface{}) (interface{}, error) {
	var arg interface{}
	forceExpression := false
	if !isFunctionBody(expression) {
		forceExpression = true
	}
	if len(options) == 1 {
		arg = options[0]
	} else if len(options) == 2 {
		arg = options[0]
		forceExpression = options[1].(bool)
	}
	result, err := b.channel.Send("evaluateExpression", map[string]interface{}{
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (f *Frame) EvaluateOnSelector(selector string, expression string, options ...interface{}) (interface{}, error) {
	var arg interface{}
	forceExpression := false
	if !isFunctionBody(expression) {
		forceExpression = true
	}
	if len(options) == 1 {
		arg = options[0]
	} else if len(options) == 2 {
		arg = options[0]
		forceExpression = options[1].(bool)
	}
	result, err := f.channel.Send("evalOnSelector", map[string]interface{}{
		"selector":   selector,
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (f *Frame) EvaluateOnSelectorAll(selector string, expression string, options ...interface{}) (interface{}, error) {
	var arg interface{}
	forceExpression := false
	if !isFunctionBody(expression) {
		forceExpression = true
	}
	if len(options) == 1 {
		arg = options[0]
	} else if len(options) == 2 {
		arg = options[0]
		forceExpression = options[1].(bool)
	}
	result, err := f.channel.Send("evalOnSelectorAll", map[string]interface{}{
		"selector":   selector,
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (f *Frame) Click(selector string, options ...PageClickOptions) error {
	_, err := f.channel.Send("click", map[string]interface{}{
		"selector": selector,
	}, options)
	return err
}

func newFrame(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Frame {
	bt := &Frame{
		url: initializer["url"].(string),
	}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("navigated", bt.onFrameNavigated)
	return bt
}
