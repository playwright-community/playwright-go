package playwright

type Worker struct {
	ChannelOwner
}

func newWorker(parent *ChannelOwner, objectType string, guid string, initializer map[string]interface{}) *Worker {
	bt := &Worker{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
