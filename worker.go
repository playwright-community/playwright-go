package playwright

type Worker struct {
	ChannelOwner
	page *Page
}

func (w *Worker) URL() string {
	return w.initializer["url"].(string)
}

func (b *Worker) Evaluate(expression string, options ...interface{}) (interface{}, error) {
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

func newWorker(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Worker {
	bt := &Worker{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("close", func(payload ...interface{}) {
		workers := make([]*Worker, 0)
		for i := 0; i < len(bt.page.workers); i++ {
			if bt.page.workers[i] != bt {
				workers = append(workers, bt.page.workers[i])
			}
		}
		bt.page.workers = workers
		bt.Emit("close", bt)
	})
	return bt
}
