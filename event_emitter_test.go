package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testEventName    = "foobar"
	testEventNameFoo = "foo"
	testEventNameBar = "bar"
)

func TestEventEmitterListenerCount(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan interface{}, 1)
	myHandler := func(payload ...interface{}) {
		wasCalled <- payload[0]
	}
	require.Nil(t, handler.events[testEventNameFoo])
	handler.On(testEventNameFoo, myHandler)
	require.Equal(t, 1, handler.ListenerCount(testEventNameFoo))
	handler.Once(testEventNameFoo, myHandler)
	require.Equal(t, 2, handler.ListenerCount(testEventNameFoo))
	require.Nil(t, handler.events[testEventNameBar])
	handler.Once(testEventNameBar, myHandler)
	require.Equal(t, 1, handler.ListenerCount(testEventNameBar))
	require.Equal(t, 3, handler.ListenerCount(""))
}

func TestEventEmitterOn(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan interface{}, 1)
	require.Nil(t, handler.events[testEventName])
	handler.On(testEventName, func(payload ...interface{}) {
		wasCalled <- payload[0]
	})
	require.Equal(t, 1, handler.ListenerCount(testEventName))
	value := 123
	handler.Emit(testEventName, value)
	result := <-wasCalled
	require.Equal(t, 1, handler.ListenerCount(testEventName))
	require.Equal(t, result.(int), value)
}

func TestEventEmitterOnce(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan interface{}, 1)
	require.Nil(t, handler.events[testEventName])
	handler.Once(testEventName, func(payload ...interface{}) {
		wasCalled <- payload[0]
	})
	require.Equal(t, 1, handler.ListenerCount(testEventName))
	value := 123
	handler.Emit(testEventName, value)
	result := <-wasCalled
	require.Equal(t, result.(int), value)
	require.Equal(t, 0, handler.ListenerCount(testEventName))
}

func TestEventEmitterRemove(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan interface{}, 1)
	require.Nil(t, handler.events[testEventName])
	myHandler := func(payload ...interface{}) {
		wasCalled <- payload[0]
	}
	handler.On(testEventName, myHandler)
	require.Equal(t, 1, handler.ListenerCount(testEventName))
	value := 123
	handler.Emit(testEventName, value)
	result := <-wasCalled
	require.Equal(t, 1, handler.ListenerCount(testEventName))
	require.Equal(t, result.(int), value)
	handler.Once(testEventName, myHandler)
	handler.RemoveListener(testEventName, myHandler)
	require.Equal(t, 0, handler.ListenerCount(testEventName))
}

func TestEventEmitterRemoveEmpty(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	handler.RemoveListener(testEventName, func(...interface{}) {})
	require.Equal(t, 0, handler.ListenerCount(testEventName))
}

func TestEventEmitterRemoveKeepExisting(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	handler.On(testEventName, func(...interface{}) {})
	handler.Once(testEventName, func(...interface{}) {})
	handler.RemoveListener("abc123", func(...interface{}) {})
	handler.RemoveListener(testEventName, func(...interface{}) {})
	require.Equal(t, 2, handler.ListenerCount(testEventName))
}

func TestEventEmitterOnLessArgsAcceptingReceiver(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan bool, 1)
	require.Nil(t, handler.events[testEventName])
	handler.Once(testEventName, func(ev ...interface{}) {
		wasCalled <- true
	})
	handler.Emit(testEventName)
	<-wasCalled
}
