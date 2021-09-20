package playwright

import (
	"sync"
)

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

func (c *channelOwner) dispose() {
	// Clean up from parent and connection.
	if c.parent != nil {
		delete(c.parent.objects, c.guid)
	}
	delete(c.connection.objects, c.guid)

	// Dispose all children.
	for _, object := range c.objects {
		object.dispose()
	}
	c.objects = make(map[string]*channelOwner)
}

func (c *channelOwner) createChannelOwner(self interface{}, parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) {
	c.objectType = objectType
	c.guid = guid
	c.parent = parent
	c.objects = make(map[string]*channelOwner)
	c.initializer = initializer
	if c.parent != nil {
		c.connection = parent.connection
		c.parent.objects[guid] = c
	}
	if c.connection != nil {
		c.connection.objects[guid] = c
	}
	c.channel = newChannel(c.connection, guid)
	c.channel.object = self
	c.initEventEmitter()
}

type rootChannelOwner struct {
	channelOwner
}

func (r *rootChannelOwner) initialize() (*Playwright, error) {
	result, err := r.channel.Send("initialize", map[string]interface{}{
		"sdkLanguage": "javascript",
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(result).(*Playwright), nil
}

func newRootChannelOwner(connection *connection) *rootChannelOwner {
	c := &rootChannelOwner{}
	c.connection = connection
	c.createChannelOwner(c, nil, "Root", "", make(map[string]interface{}))
	return c
}
