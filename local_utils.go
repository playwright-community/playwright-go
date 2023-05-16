package playwright

type localUtilsImpl struct {
	channelOwner
}

type localUtilsZipOptions struct {
	ZipFile        string        `json:"zipFile"`
	Entries        []interface{} `json:"entries"`
	StacksId       string        `json:"stacksId"`
	Mode           string        `json:"mode"`
	IncludeSources bool          `json:"includeSources"`
}

func (l *localUtilsImpl) Zip(options localUtilsZipOptions) (interface{}, error) {
	return l.channel.Send("zip", options)
}

func (l *localUtilsImpl) TracingStarted(traceName string, traceDir ...string) (string, error) {
	overrides := make(map[string]interface{})
	overrides["traceName"] = traceName
	if len(traceDir) > 0 {
		overrides["traceDir"] = traceDir[0]
	}
	stacksId, err := l.channel.Send("tracingStarted", overrides)
	if stacksId == nil {
		return "", err
	}
	return stacksId.(string), err
}

func (l *localUtilsImpl) TraceDiscarded(stacksId string) error {
	_, err := l.channel.Send("traceDiscarded", map[string]interface{}{
		"stacksId": stacksId,
	})
	return err
}

func (l *localUtilsImpl) AddStackToTracingNoReply(id int, stack []map[string]interface{}) {
	l.channel.SendNoReply("addStackToTracingNoReply", map[string]interface{}{
		"callData": map[string]interface{}{
			"id":    id,
			"stack": stack,
		},
	})
}

func newLocalUtils(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *localUtilsImpl {
	l := &localUtilsImpl{}
	l.createChannelOwner(l, parent, objectType, guid, initializer)
	return l
}
