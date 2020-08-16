package playwright

type ChannelOwner struct {
	objectType  string
	guid        string
	channel     *Channel
	objects     map[string]*ChannelOwner
	connection  *Connection
	initializer interface{}
	parent      *ChannelOwner
}

func (c *ChannelOwner) Dispose() error {
	return nil
}

func (c *ChannelOwner) createChannelOwner(self interface{}, parent *ChannelOwner, objectType string, guid string, initializer interface{}) {
	c.objectType = objectType
	c.guid = guid
	c.parent = parent
	c.objects = make(map[string]*ChannelOwner)
	c.connection = parent.connection
	c.channel = newChannel(c.connection, guid)
	c.channel.object = self
	c.initializer = initializer
	c.connection.objects[guid] = c
	c.parent.objects[guid] = c
}

func newRootChannelOwner(connection *Connection) *ChannelOwner {
	c := &ChannelOwner{
		objectType: "",
		guid:       "",
		connection: connection,
		objects:    make(map[string]*ChannelOwner),
		channel:    newChannel(connection, ""),
	}
	c.channel.object = c
	c.connection.objects[""] = c
	return c
}
