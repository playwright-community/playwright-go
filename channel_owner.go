package playwright

import (
	"sync"
)

type channelOwner struct {
	sync.RWMutex
	eventEmitter
	objectType                 string
	guid                       string
	channel                    *channel
	objects                    map[string]*channelOwner
	eventToSubscriptionMapping map[string]string
	connection                 *connection
	initializer                map[string]interface{}
	parent                     *channelOwner
	wasCollected               bool
	isInternalType             bool
}

func (c *channelOwner) dispose(reason ...string) {
	// Clean up from parent and connection.
	if c.parent != nil {
		delete(c.parent.objects, c.guid)
	}
	c.connection.objects.Delete(c.guid)
	if len(reason) > 0 {
		c.wasCollected = reason[0] == "gc"
	}

	// Dispose all children.
	for _, object := range c.objects {
		object.dispose(reason...)
	}
	c.objects = make(map[string]*channelOwner)
}

func (c *channelOwner) adopt(child *channelOwner) {
	delete(child.parent.objects, child.guid)
	c.objects[child.guid] = child
	child.parent = c
}

func (c *channelOwner) setEventSubscriptionMapping(mapping map[string]string) {
	c.eventToSubscriptionMapping = mapping
}

func (c *channelOwner) updateSubscription(event string, enabled bool) {
	protocolEvent, ok := c.eventToSubscriptionMapping[event]
	if ok {
		c.channel.SendNoReplyInternal("updateSubscription", map[string]interface{}{
			"event":   protocolEvent,
			"enabled": enabled,
		})
	}
}

func (c *channelOwner) Once(name string, handler interface{}) {
	c.addEvent(name, handler, true)
}

func (c *channelOwner) On(name string, handler interface{}) {
	c.addEvent(name, handler, false)
}

func (c *channelOwner) addEvent(name string, handler interface{}, once bool) {
	if c.ListenerCount(name) == 0 {
		c.updateSubscription(name, true)
	}
	c.eventEmitter.addEvent(name, handler, once)
}

func (c *channelOwner) RemoveListener(name string, handler interface{}) {
	c.eventEmitter.RemoveListener(name, handler)
	if c.ListenerCount(name) == 0 {
		c.updateSubscription(name, false)
	}
}

func (c *channelOwner) createChannelOwner(self interface{}, parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) {
	c.objectType = objectType
	c.guid = guid
	c.wasCollected = false
	c.parent = parent
	c.objects = make(map[string]*channelOwner)
	c.initializer = initializer
	if c.parent != nil {
		c.connection = parent.connection
		c.parent.objects[guid] = c
	}
	if c.connection != nil {
		c.connection.objects.Store(guid, c)
	}
	c.channel = newChannel(c, self)
	c.eventToSubscriptionMapping = map[string]string{}
}

func (c *channelOwner) markAsInternalType() {
	c.isInternalType = true
}

type rootChannelOwner struct {
	channelOwner
}

func (r *rootChannelOwner) initialize() (*Playwright, error) {
	ret, err := r.channel.SendReturnAsDict("initialize", map[string]interface{}{
		"sdkLanguage": "javascript",
	})
	if err != nil {
		return nil, err
	}
	return fromChannel(ret["playwright"]).(*Playwright), nil
}

func newRootChannelOwner(connection *connection) *rootChannelOwner {
	c := &rootChannelOwner{}
	c.connection = connection
	c.createChannelOwner(c, nil, "Root", "", make(map[string]interface{}))
	return c
}
