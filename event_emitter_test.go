package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testEventName = "foobar"

func TestEventEmitterOn(t *testing.T) {
	handler := &EventEmitter{}
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
	handler := &EventEmitter{}
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
	handler := &EventEmitter{}
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
	handler := &EventEmitter{}
	handler.initEventEmitter()
	handler.RemoveListener(testEventName, func(...interface{}) {})
	require.Equal(t, 0, handler.ListenerCount(testEventName))
}
