package playwright

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	testEventNameFoobar = "foobar"
	testEventNameReject = "reject"
	testEventPayload    = "payload data"
)

func TestWaiterWaitForEvent(t *testing.T) {
	timeout := 500.0
	emitter := &eventEmitter{}
	emitter.initEventEmitter()
	waiter := newWaiter()
	waiter.RejectOnTimeout(timeout)
	evChan := waiter.WaitForEvent(emitter, testEventNameFoobar)
	go func() {
		emitter.Emit(testEventNameFoobar, testEventPayload)
		emitter.Emit(testEventNameFoobar, "2")
	}()
	result := <-evChan
	require.NoError(t, waiter.Err())
	require.Equal(t, result, testEventPayload)
}

func TestWaiterRejectOnTimeout(t *testing.T) {
	timeout := 500.0
	emitter := &eventEmitter{}
	emitter.initEventEmitter()
	waiter := newWaiter()
	waiter.RejectOnTimeout(timeout)
	evChan := waiter.WaitForEvent(emitter, testEventNameFoobar)
	go func() {
		time.Sleep(time.Duration(timeout+3) * time.Millisecond)
		emitter.Emit(testEventNameFoobar, testEventPayload)
	}()
	result := <-evChan
	require.EqualError(t, waiter.Err(), fmt.Sprintf("Timeout %.2fms exceeded.", timeout))
	require.Nil(t, result)
}

func TestWaiterRejectOnEvent(t *testing.T) {
	errCause := fmt.Errorf("reject on event")
	emitter := &eventEmitter{}
	emitter.initEventEmitter()
	waiter := newWaiter()
	waiter.RejectOnEvent(emitter, testEventNameReject, errCause)
	evChan := waiter.WaitForEvent(emitter, testEventNameFoobar)
	require.Equal(t, 1, emitter.ListenerCount(testEventNameReject))
	go func() {
		emitter.Emit(testEventNameReject)
		emitter.Emit(testEventNameFoobar, testEventPayload)
		emitter.Emit(testEventNameFoobar, "3")
	}()
	result := <-evChan
	require.EqualError(t, waiter.Err(), errCause.Error())
	require.Equal(t, 0, emitter.ListenerCount(testEventNameReject))
	require.Nil(t, result)
}
