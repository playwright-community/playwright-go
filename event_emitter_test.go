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
	require.Equal(t, 1, len(handler.events[testEventName].events))
	value := 123
	handler.Emit(testEventName, value)
	result := <-wasCalled
	require.Equal(t, 1, len(handler.events[testEventName].events))
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
	require.Equal(t, 1, len(handler.events[testEventName].events))
	value := 123
	handler.Emit(testEventName, value)
	result := <-wasCalled
	require.Equal(t, result.(int), value)
	require.Equal(t, 0, len(handler.events[testEventName].events))
}
