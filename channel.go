package playwright

import (
	"encoding/json"
	"fmt"
)

type channel struct {
	eventEmitter
	guid       string
	connection *connection
	owner      *channelOwner // to avoid type conversion
	object     interface{}   // retain type info (for fromChannel needed)
}

func (c *channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"guid": c.guid,
	})
}

// for catch errors of route handlers etc.
func (c *channel) CreateTask(fn func()) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				err, ok := e.(error)
				if ok {
					c.connection.err.Set(err)
				} else {
					c.connection.err.Set(fmt.Errorf("%v", e))
				}
			}
		}()
		fn()
	}()
}

func (c *channel) Send(method string, options ...interface{}) (interface{}, error) {
	return c.connection.WrapAPICall(func() (interface{}, error) {
		return c.innerSend(method, options...).GetResultValue()
	}, c.owner.isInternalType)
}

func (c *channel) SendReturnAsDict(method string, options ...interface{}) (map[string]interface{}, error) {
	ret, err := c.connection.WrapAPICall(func() (interface{}, error) {
		return c.innerSend(method, options...).GetResult()
	}, c.owner.isInternalType)
	return ret.(map[string]interface{}), err
}

func (c *channel) innerSend(method string, options ...interface{}) *protocolCallback {
	if err := c.connection.err.Get(); err != nil {
		c.connection.err.Set(nil)
		pc := newProtocolCallback(false, c.connection.abort)
		pc.SetError(err)
		return pc
	}
	params := transformOptions(options...)
	return c.connection.sendMessageToServer(c.owner, method, params, false)
}

// SendNoReply ignores return value and errors
// almost equivalent to `send(...).catch(() => {})`
func (c *channel) SendNoReply(method string, options ...interface{}) {
	c.innerSendNoReply(method, c.owner.isInternalType, options...)
}

func (c *channel) SendNoReplyInternal(method string, options ...interface{}) {
	c.innerSendNoReply(method, true, options...)
}

func (c *channel) innerSendNoReply(method string, isInternal bool, options ...interface{}) {
	params := transformOptions(options...)
	_, err := c.connection.WrapAPICall(func() (interface{}, error) {
		return c.connection.sendMessageToServer(c.owner, method, params, true).GetResult()
	}, isInternal)
	if err != nil {
		// ignore error actively, log only for debug
		logger.Error("SendNoReply failed", "error", err)
	}
}

func newChannel(owner *channelOwner, object interface{}) *channel {
	channel := &channel{
		connection: owner.connection,
		guid:       owner.guid,
		owner:      owner,
		object:     object,
	}
	return channel
}
