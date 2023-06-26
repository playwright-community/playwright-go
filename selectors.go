package playwright

import (
	"errors"
	"os"
	"sync"
)

type selectorsOwnerImpl struct {
	channelOwner
}

func newSelectorsOwner(parent *channelOwner, objectType string, guid string, initializer map[string]interface{}) *selectorsOwnerImpl {
	obj := &selectorsOwnerImpl{}
	obj.createChannelOwner(obj, parent, objectType, guid, initializer)
	return obj
}

type selectorsImpl struct {
	channels      sync.Map
	registrations []map[string]interface{}
}

func (s *selectorsImpl) Register(name string, option SelectorsRegisterOptions) error {
	if option.Script == nil && option.Path == nil {
		return errors.New("Either source or path should be specified")
	}
	script := ""
	if option.Path != nil {
		content, err := os.ReadFile(*option.Path)
		if err != nil {
			return err
		}
		script = string(content)
	} else {
		script = *option.Script
	}
	params := map[string]interface{}{
		"name":   name,
		"source": script,
	}
	if option.ContentScript != nil && *option.ContentScript {
		params["contentScript"] = true
	}
	var err error
	s.channels.Range(func(key, value any) bool {
		_, err = value.(*selectorsOwnerImpl).channel.Send("register", params)
		return err == nil
	})
	if err != nil {
		return err
	}
	s.registrations = append(s.registrations, params)
	return nil
}

func (s *selectorsImpl) SetTestIdAttribute(name string) {
	setTestIdAttributeName(name)
	s.channels.Range(func(key, value any) bool {
		value.(*selectorsOwnerImpl).channel.SendNoReply("setTestIdAttributeName", map[string]interface{}{
			"testIdAttributeName": name,
		})
		return true
	})
}

func (s *selectorsImpl) addChannel(channel *selectorsOwnerImpl) {
	s.channels.Store(channel.guid, channel)
	for _, params := range s.registrations {
		channel.channel.SendNoReply("register", params)
		channel.channel.SendNoReply("setTestIdAttributeName", map[string]interface{}{
			"testIdAttributeName": getTestIdAttributeName(),
		})
	}
}

func (s *selectorsImpl) removeChannel(channel *selectorsOwnerImpl) {
	s.channels.Delete(channel.guid)
}

func newSelectorsImpl() *selectorsImpl {
	return &selectorsImpl{
		channels:      sync.Map{},
		registrations: make([]map[string]interface{}, 0),
	}
}
