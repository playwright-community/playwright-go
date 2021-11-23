package playwright

type tracingImpl struct {
	context *browserContextImpl
	channel *channel
}

func (t *tracingImpl) Start(options ...TracingStartOptions) error {
	if _, err := t.channel.Send("tracingStart", options); err != nil {
		return err
	}
	if _, err := t.channel.Send("tracingStartChunk", options); err != nil {
		return err
	}
	return nil
}

func (t *tracingImpl) StartChunk() error {
	_, err := t.channel.Send("tracingStartChunk")
	return err
}
func (t *tracingImpl) Stop(options ...TracingStopOptions) error {
	path := ""
	if len(options) == 1 && options[0].Path != nil {
		path = *options[0].Path
	}
	if err := t.stopChunk(path); err != nil {
		return err
	}
	if _, err := t.channel.Send("tracingStop"); err != nil {
		return err
	}
	return nil
}
func (t *tracingImpl) StopChunk(options ...TracingStopChunkOptions) error {
	path := ""
	if len(options) == 1 && options[0].Path != nil {
		path = *options[0].Path
	}
	if err := t.stopChunk(path); err != nil {
		return err
	}
	return nil
}

func (t *tracingImpl) stopChunk(path string) error {
	save := true
	if path == "" {
		save = false
	}
	artifactChannel, err := t.channel.Send("tracingStopChunk", map[string]interface{}{
		"save":         save,
		"skipCompress": false,
	})
	if err != nil {
		return err
	}
	if artifactChannel != nil {
		artifact := fromChannel(artifactChannel).(*artifactImpl)
		if path != "" {
			if err := artifact.SaveAs(path); err != nil {
				return err
			}
		}
		return artifact.Delete()
	}
	return nil
}

func newTracing(context *browserContextImpl) *tracingImpl {
	return &tracingImpl{context, context.channel}
}
