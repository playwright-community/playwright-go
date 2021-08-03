package playwright

import "sync"

type channelOwner struct {
	sync.RWMutex
	eventEmitter
	objectType  string
	guid        string
	channel     *channel
	objects     map[string]*channelOwner
	connection  *connection
	initializer map[string]interface{}
	parent      *channelOwner
}

func (c *channelOwner) Dispose() {
	// Clean up from parent and connection.
	if c.parent != nil {
		delete(c.parent.objects, c.guid)
	}
	delete(c.connection.objects, c.guid)

	// Dispose all children.
	for _, object := range c.objects {
		object.Dispose()
	}
	c.objects = make(map[string]*channelOwner)
}

func (c *channelOwner) createChannelOwner(self interface{}, parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) {
	c.objectType = objectType
	c.guid = guid
	c.parent = parent
	c.objects = make(map[string]*channelOwner)
	c.connection = parent.connection
	c.channel = newChannel(c.connection, guid)
	c.channel.object = self
	c.initializer = initializer
	c.connection.objects[guid] = c
	c.parent.objects[guid] = c
	c.initEventEmitter()
}

func newRootChannelOwner(connection *connection) *channelOwner {
	c := &channelOwner{
		objectType: "",
		guid:       "",
		connection: connection,
		objects:    make(map[string]*channelOwner),
		channel:    newChannel(connection, ""),
	}
	c.channel.object = c
	c.connection.objects[""] = c
	c.initEventEmitter()
	return c
}
