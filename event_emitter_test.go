package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testEventName = "foobar"

func TestEventEmitterOn(t *testing.T) {
	handler := &EventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan interface{})
	require.Nil(t, handler.events[testEventName])
	handler.On(testEventName, func(payload ...interface{}) {
		wasCalled <- payload[0]
	})
	require.Equal(t, 1, len(handler.events[testEventName].on))
	require.Equal(t, 0, len(handler.events[testEventName].once))
	value := 123
	handler.Emit(testEventName, value)
	result := <-wasCalled
	require.Equal(t, 1, len(handler.events[testEventName].on))
	require.Equal(t, 0, len(handler.events[testEventName].once))
	require.Equal(t, result.(int), value)
}

func TestEventEmitterOnce(t *testing.T) {
	handler := &EventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan interface{})
	require.Nil(t, handler.events[testEventName])
	handler.Once(testEventName, func(payload ...interface{}) {
		wasCalled <- payload[0]
	})
	require.Equal(t, 0, len(handler.events[testEventName].on))
	require.Equal(t, 1, len(handler.events[testEventName].once))
	value := 123
	handler.Emit(testEventName, value)
	result := <-wasCalled
	require.Equal(t, result.(int), value)
	require.Equal(t, 0, len(handler.events[testEventName].on))
	require.Equal(t, 0, len(handler.events[testEventName].once))
}

func TestEventEmitterRemove(t *testing.T) {
	handler := &EventEmitter{}
	handler.initEventEmitter()
	wasCalled := make(chan interface{})
	require.Nil(t, handler.events[testEventName])
	myHandler := func(payload ...interface{}) {
		wasCalled <- payload[0]
	}
	handler.On(testEventName, myHandler)
	require.Equal(t, 1, len(handler.events[testEventName].on))
	require.Equal(t, 0, len(handler.events[testEventName].once))
	value := 123
	handler.Emit(testEventName, value)
	result := <-wasCalled
	require.Equal(t, 1, len(handler.events[testEventName].on))
	require.Equal(t, 0, len(handler.events[testEventName].once))
	require.Equal(t, result.(int), value)
	handler.RemoveListener(testEventName, myHandler)
	require.Equal(t, 1, len(handler.events[testEventName].on))
	require.Equal(t, 0, len(handler.events[testEventName].once))
}

func TestEventEmitterRemoveEmpty(t *testing.T) {
	handler := &EventEmitter{}
	handler.initEventEmitter()
	handler.RemoveListener(testEventName, func(...interface{}) {})
}
