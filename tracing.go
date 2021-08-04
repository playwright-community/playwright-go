package playwright

type tracingImpl struct {
	context *browserContextImpl
	channel *channel
}

func (t *tracingImpl) Start(options ...TracingStartOptions) error {
	_, err := t.channel.Send("tracingStart", options)
	return err
}

func (t *tracingImpl) Stop(options ...TracingStopOptions) error {
	if _, err := t.channel.Send("tracingStop", nil); err != nil {
		return err
	}
	if len(options) == 1 && options[0].Path != nil {
		artifactChannel, err := t.channel.Send("tracingExport", nil)
		if err != nil {
			return err
		}
		artifact := fromChannel(artifactChannel).(*artifactImpl)
		if err = artifact.SaveAs(*options[0].Path); err != nil {
			return err
		}
		if err = artifact.Delete(); err != nil {
			return err
		}
	}
	return nil
}

func newTracing(context *browserContextImpl) *tracingImpl {
	return &tracingImpl{context, context.channel}
}
