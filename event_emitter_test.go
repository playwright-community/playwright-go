package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
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
	require.Nil(t, handler.events[testEventNameFoo])
	handler.On(testEventNameFoo, func(payload ...interface{}) {
		wasCalled <- payload[0]
	})
	require.Equal(t, 1, handler.ListenerCount(testEventNameFoo))
	value := 123
	handler.Emit(testEventNameFoo, value)
	result := <-wasCalled
	require.Equal(t, 1, handler.ListenerCount(testEventNameFoo))
	require.Equal(t, result.(int), value)
}

func TestEventEmitterOnce(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan interface{}, 1)
	require.Nil(t, handler.events[testEventNameFoo])
	handler.Once(testEventNameFoo, func(payload ...interface{}) {
		wasCalled <- payload[0]
	})
	require.Equal(t, 1, handler.ListenerCount(testEventNameFoo))
	value := 123
	handler.Emit(testEventNameFoo, value)
	result := <-wasCalled
	require.Equal(t, result.(int), value)
	require.Equal(t, 0, handler.ListenerCount(testEventNameFoo))
}

func TestEventEmitterRemove(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan interface{}, 1)
	require.Nil(t, handler.events[testEventNameFoo])
	myHandler := func(payload ...interface{}) {
		wasCalled <- payload[0]
	}
	handler.On(testEventNameFoo, myHandler)
	require.Equal(t, 1, handler.ListenerCount(testEventNameFoo))
	value := 123
	handler.Emit(testEventNameFoo, value)
	result := <-wasCalled
	require.Equal(t, 1, handler.ListenerCount(testEventNameFoo))
	require.Equal(t, result.(int), value)
	handler.Once(testEventNameFoo, myHandler)
	handler.RemoveListener(testEventNameFoo, myHandler)
	require.Equal(t, 0, handler.ListenerCount(testEventNameFoo))
}

func TestEventEmitterRemoveEmpty(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	handler.RemoveListener(testEventNameFoo, func(...interface{}) {})
	require.Equal(t, 0, handler.ListenerCount(testEventNameFoo))
}

func TestEventEmitterRemoveKeepExisting(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	handler.On(testEventNameFoo, func(...interface{}) {})
	handler.Once(testEventNameFoo, func(...interface{}) {})
	handler.RemoveListener("abc123", func(...interface{}) {})
	handler.RemoveListener(testEventNameFoo, func(...interface{}) {})
	require.Equal(t, 2, handler.ListenerCount(testEventNameFoo))
}

func TestEventEmitterOnLessArgsAcceptingReceiver(t *testing.T) {
	handler := &eventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan bool, 1)
	require.Nil(t, handler.events[testEventNameFoo])
	handler.Once(testEventNameFoo, func(ev ...interface{}) {
		wasCalled <- true
	})
	handler.Emit(testEventNameFoo)
	<-wasCalled
}
