package playwright

import "fmt"

type tracingImpl struct {
	channelOwner
	includeSources bool
	isTracing      bool
	stacksId       string
	tracesDir      string
}

func (t *tracingImpl) Start(options ...TracingStartOptions) error {
	chunkOption := TracingStartChunkOptions{}
	if len(options) == 1 {
		if options[0].Sources != nil {
			t.includeSources = *options[0].Sources
		}
		chunkOption.Name = options[0].Name
		chunkOption.Title = options[0].Title
	}
	innerStart := func() (interface{}, error) {
		if _, err := t.channel.Send("tracingStart", options); err != nil {
			return "", err
		}
		return t.channel.Send("tracingStartChunk", options)
	}
	result, err := t.connection.WrapAPICall(innerStart, true)
	if err != nil {
		return err
	}
	name, ok := result.(string)
	if !ok {
		name = ""
	}
	return t.startCollectingStacks(name)
}

func (t *tracingImpl) StartChunk(options ...TracingStartChunkOptions) error {
	result, err := t.channel.Send("tracingStartChunk", options)
	if err != nil {
		return err
	}
	name, ok := result.(string)
	if !ok {
		name = ""
	}
	return t.startCollectingStacks(name)
}

func (t *tracingImpl) StopChunk(path ...string) error {
	filePath := ""
	if len(path) == 1 {
		filePath = path[0]
	}
	_, err := t.connection.WrapAPICall(func() (interface{}, error) {
		return nil, t.doStopChunk(filePath)
	}, true)
	return err
}

func (t *tracingImpl) Stop(path ...string) error {
	filePath := ""
	if len(path) == 1 {
		filePath = path[0]
	}
	_, err := t.connection.WrapAPICall(func() (interface{}, error) {
		if err := t.doStopChunk(filePath); err != nil {
			return nil, err
		}
		return t.channel.Send("tracingStop")
	}, true)
	return err
}

func (t *tracingImpl) doStopChunk(filePath string) (err error) {
	if t.isTracing {
		t.isTracing = false
		t.connection.setInTracing(false)
	}
	if filePath == "" {
		// Not interested in artifacts.
		_, err = t.channel.Send("tracingStopChunk", map[string]interface{}{
			"mode": "discard",
		})
		if t.stacksId != "" {
			return t.connection.LocalUtils().TraceDiscarded(t.stacksId)
		}
		return err
	}

	isLocal := !t.connection.isRemote
	if isLocal {
		result, err := t.channel.SendReturnAsDict("tracingStopChunk", map[string]interface{}{
			"mode": "entries",
		})
		if err != nil {
			return err
		}
		entries, ok := result.(map[string]interface{})["entries"]
		if !ok {
			return fmt.Errorf("could not convert result to map: %v", result)
		}
		_, err = t.connection.LocalUtils().Zip(localUtilsZipOptions{
			ZipFile:        filePath,
			Entries:        entries.([]interface{}),
			StacksId:       t.stacksId,
			Mode:           "write",
			IncludeSources: t.includeSources,
		})
		return err
	}

	result, err := t.channel.SendReturnAsDict("tracingStopChunk", map[string]interface{}{
		"mode": "archive",
	})
	if err != nil {
		return err
	}
	artifactChannel, ok := result.(map[string]interface{})["artifact"]
	if !ok {
		return fmt.Errorf("could not convert result to map: %v", result)
	}
	// Save trace to the final local file.
	artifact := fromNullableChannel(artifactChannel).(*artifactImpl)
	// The artifact may be missing if the browser closed while stopping tracing.
	if artifact == nil {
		if t.stacksId != "" {
			return t.connection.LocalUtils().TraceDiscarded(t.stacksId)
		}
		return
	}
	if err := artifact.SaveAs(filePath); err != nil {
		return err
	}
	if err := artifact.Delete(); err != nil {
		return err
	}
	_, err = t.connection.LocalUtils().Zip(localUtilsZipOptions{
		ZipFile:        filePath,
		Entries:        []interface{}{},
		StacksId:       t.stacksId,
		Mode:           "append",
		IncludeSources: t.includeSources,
	})
	return err
}

func (t *tracingImpl) startCollectingStacks(name string) (err error) {
	if !t.isTracing {
		t.isTracing = true
		t.connection.setInTracing(true)
	}
	t.stacksId, err = t.connection.LocalUtils().TracingStarted(name, t.tracesDir)
	return
}

func newTracing(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *tracingImpl {
	bt := &tracingImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
