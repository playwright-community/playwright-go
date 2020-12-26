package playwright

type workerImpl struct {
	channelOwner
	page *pageImpl
}

func (w *workerImpl) URL() string {
	return w.initializer["url"].(string)
}

func (w *workerImpl) Evaluate(expression string, options ...interface{}) (interface{}, error) {
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
	result, err := w.channel.Send("evaluateExpression", map[string]interface{}{
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return parseResult(result), nil
}

func (w *workerImpl) EvaluateHandle(expression string, options ...interface{}) (JSHandle, error) {
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
	result, err := w.channel.Send("evaluateExpressionHandle", map[string]interface{}{
		"expression": expression,
		"isFunction": !forceExpression,
		"arg":        serializeArgument(arg),
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(result).(*jsHandleImpl), nil
}

func newWorker(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *workerImpl {
	bt := &workerImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	bt.channel.On("close", func() {
		workers := make([]Worker, 0)
		if bt.page != nil {
			for i := 0; i < len(bt.page.workers); i++ {
				if bt.page.workers[i].(*workerImpl) != bt {
					workers = append(workers, bt.page.workers[i])
				}
			}
			bt.page.workersLock.Lock()
			bt.page.workers = workers
			bt.page.workersLock.Unlock()
		}
		bt.Emit("close", bt)
	})
	return bt
}
