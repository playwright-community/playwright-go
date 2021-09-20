package playwright

import (
	"log"
)

type bindingCallImpl struct {
	channelOwner
}

// BindingSource is the value passed to a binding call execution
type BindingSource struct {
	Context BrowserContext
	Page    Page
	Frame   Frame
}

// ExposedFunction represents the func signature of an exposed function
type ExposedFunction = func(args ...interface{}) interface{}

// BindingCallFunction represents the func signature of an exposed binding call func
type BindingCallFunction func(source *BindingSource, args ...interface{}) interface{}

func (b *bindingCallImpl) Call(f BindingCallFunction) {
	defer func() {
		if r := recover(); r != nil {
			if _, err := b.channel.Send("reject", map[string]interface{}{
				"error": serializeError(r.(error)),
			}); err != nil {
				log.Printf("could not reject BindingCall: %v", err)
			}
		}
	}()

	frame := fromChannel(b.initializer["frame"]).(*frameImpl)
	source := &BindingSource{
		Context: frame.Page().Context(),
		Page:    frame.Page(),
		Frame:   frame,
	}
	var result interface{}
	if handle, ok := b.initializer["handle"]; ok {
		result = f(source, fromChannel(handle))
	} else {
		initializerArgs := b.initializer["args"].([]interface{})
		funcArgs := []interface{}{}
		for i := 0; i < len(initializerArgs); i++ {
			funcArgs = append(funcArgs, parseResult(initializerArgs[i]))
		}
		result = f(source, funcArgs...)
	}
	_, err := b.channel.Send("resolve", map[string]interface{}{
		"result": serializeArgument(result),
	})
	if err != nil {
		log.Printf("could not resolve BindingCall: %v", err)
	}
}

func newBindingCall(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *bindingCallImpl {
	bt := &bindingCallImpl{}
	bt.createChannelOwner(bt, parent, objectType, guid, initializer)
	return bt
}
