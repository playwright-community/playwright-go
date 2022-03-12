package playwright

type tracingImpl struct {
	channelOwner
}

func (t *tracingImpl) Start(options ...TracingStartOptions) error {
	if _, err := t.channel.Send("tracingStart", options); err != nil {
		return err
	}
	title := ""
	if len(options) == 1 && options[0].Title != nil {
		title = *options[0].Title
	}
	if _, err := t.channel.Send("tracingStartChunk", map[string]interface{}{
		"title": title,
	}); err != nil {
		return err
	}
	return nil
}

func (t *tracingImpl) StartChunk(options ...TracingStartChunkOptions) error {
	_, err := t.channel.Send("tracingStartChunk", options)
	return err
}

func (t *tracingImpl) StopChunk(options ...TracingStopChunkOptions) error {
	path := ""
	if len(options) == 1 && options[0].Path != nil {
		path = *options[0].Path
	}
	if err := t.doStopChunk(path); err != nil {
		return err
	}
	return nil
}

func (t *tracingImpl) Stop(options ...TracingStopOptions) error {
	path := ""
	if len(options) == 1 && options[0].Path != nil {
		path = *options[0].Path
		return t.doStopChunk(path)
	}
	_, err := t.channel.Send("tracingStopChunk", options)
	return err
}

func (t *tracingImpl) doStopChunk(filePath string) error {
	isLocal := !t.connection.isRemote
	mode := "doNotSave"
	if filePath != "" {
		if isLocal {
			mode = "compressTraceAndSources"
		} else {
			mode = "compressTrace"
		}
	}
	artifactChannel, err := t.channel.Send("tracingStopChunk", map[string]interface{}{
		"mode": mode,
	})
	if err != nil {
		return err
	}
	// Not interested in artifacts.
	if filePath == "" {
		return nil
	}
	// The artifact may be missing if the browser closed while stopping tracing.
	if artifactChannel == nil {
		return nil
	}
	// Save trace to the final local file.
	artifact := fromChannel(artifactChannel).(*artifactImpl)
	if err := artifact.SaveAs(filePath); err != nil {
		return err
	}
	return artifact.Delete()
}

func newTracing(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *tracingImpl {
	bt := &tracingImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
