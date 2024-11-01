package playwright

import (
	"errors"
	"os"
	"sync"
)

type selectorsOwnerImpl struct {
	channelOwner
}

func (s *selectorsOwnerImpl) setTestIdAttributeName(name string) {
	s.channel.SendNoReply("setTestIdAttributeName", map[string]interface{}{
		"testIdAttributeName": name,
	})
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

func (s *selectorsImpl) Register(name string, script Script, options ...SelectorsRegisterOptions) error {
	if script.Path == nil && script.Content == nil {
		return errors.New("Either source or path should be specified")
	}
	source := ""
	if script.Path != nil {
		content, err := os.ReadFile(*script.Path)
		if err != nil {
			return err
		}
		source = string(content)
	} else {
		source = *script.Content
	}
	params := map[string]interface{}{
		"name":   name,
		"source": source,
	}
	if len(options) == 1 && options[0].ContentScript != nil {
		params["contentScript"] = *options[0].ContentScript
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
		value.(*selectorsOwnerImpl).setTestIdAttributeName(name)
		return true
	})
}

func (s *selectorsImpl) addChannel(channel *selectorsOwnerImpl) {
	s.channels.Store(channel.guid, channel)
	for _, params := range s.registrations {
		channel.channel.SendNoReply("register", params)
		channel.setTestIdAttributeName(getTestIdAttributeName())
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
