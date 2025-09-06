package playwright

import (
	"errors"
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
	waiter := newWaiter().WithTimeout(timeout)
	_, err := waiter.Wait()
	require.Error(t, err)
	waiter.WaitForEvent(emitter, testEventNameFoobar, nil)
	go func() {
		emitter.Emit(testEventNameFoobar, testEventPayload)
		emitter.Emit(testEventNameFoobar, "2")
		emitter.Emit(testEventNameFoobar, "3")
	}()
	result, err := waiter.Wait()
	require.NoError(t, err)
	require.Equal(t, result, testEventPayload)
}

func TestWaiterWaitForEventWithPredicate(t *testing.T) {
	timeout := 500.0
	emitter := &eventEmitter{}
	waiter := newWaiter().WithTimeout(timeout)
	waiter.WaitForEvent(emitter, testEventNameFoobar, func(payload interface{}) bool {
		content, ok := payload.(string)
		if ok && content == testEventPayload {
			return true
		}
		return false
	})
	go func() {
		emitter.Emit(testEventNameFoobar, "1")
		emitter.Emit(testEventNameFoobar, testEventPayload)
		emitter.Emit(testEventNameFoobar, "3")
	}()
	result, err := waiter.Wait()
	require.NoError(t, err)
	require.Equal(t, result, testEventPayload)
}

func TestWaiterRejectOnTimeout(t *testing.T) {
	timeout := 300.0
	emitter := &eventEmitter{}
	waiter := newWaiter().WithTimeout(timeout)
	waiter.WaitForEvent(emitter, testEventNameFoobar, nil)
	go func() {
		time.Sleep(time.Duration(timeout*2) * time.Millisecond)
		emitter.Emit(testEventNameFoobar, testEventPayload)
	}()
	result, err := waiter.Wait()
	require.ErrorContains(t, err, fmt.Sprintf("Timeout %.2fms exceeded.", timeout))
	require.Nil(t, result)
}

func TestWaiterRejectOnEvent(t *testing.T) {
	errCause := fmt.Errorf("reject on event")
	errPredicate := fmt.Errorf("payload on event")
	emitter := &eventEmitter{}
	waiter := newWaiter().RejectOnEvent(emitter, testEventNameReject, errCause)
	waiter.RejectOnEvent(emitter, testEventNameFoobar, errPredicate, func(payload interface{}) bool {
		content, ok := payload.(string)
		if ok && content == "testEventPayload" {
			return true
		}
		return false
	})
	waiter.WaitForEvent(emitter, testEventNameFoobar, nil)
	require.Equal(t, 1, emitter.ListenerCount(testEventNameReject))
	go func() {
		emitter.Emit(testEventNameReject)
		emitter.Emit(testEventNameFoobar, testEventPayload)
		emitter.Emit(testEventNameFoobar, "3")
	}()
	result, err := waiter.Wait()
	require.EqualError(t, err, errCause.Error())
	require.Equal(t, 0, emitter.ListenerCount(testEventNameReject))
	require.Nil(t, result)
}

func TestWaiterRejectOnEventWithPredicate(t *testing.T) {
	errCause := fmt.Errorf("reject on event")
	errPredicate := fmt.Errorf("payload on event")
	emitter := &eventEmitter{}
	waiter := newWaiter().RejectOnEvent(emitter, testEventNameReject, errCause)
	waiter.RejectOnEvent(emitter, testEventNameFoobar, errPredicate, func(payload interface{}) bool {
		content, ok := payload.(string)
		if ok && content == testEventPayload {
			return true
		}
		return false
	})
	waiter.WaitForEvent(emitter, testEventNameFoobar, nil)
	require.Equal(t, 1, emitter.ListenerCount(testEventNameReject))
	go func() {
		emitter.Emit(testEventNameFoobar, testEventPayload)
		emitter.Emit(testEventNameReject)
		emitter.Emit(testEventNameFoobar, "3")
	}()
	result, err := waiter.Wait()
	require.EqualError(t, err, errPredicate.Error())
	require.Equal(t, 0, emitter.ListenerCount(testEventNameReject))
	require.Nil(t, result)
}

func TestWaiterReturnErrorWhenMisuse(t *testing.T) {
	emitter := &eventEmitter{}
	waiter := newWaiter()
	waiter.WaitForEvent(emitter, testEventNameFoobar, nil)
	waiter.WithTimeout(500)
	_, err := waiter.Wait()
	require.ErrorContains(t, err, "please set timeout before WaitForEvent")

	waiter = newWaiter()
	waiter.WaitForEvent(emitter, testEventNameFoobar, nil)
	waiter.WaitForEvent(emitter, testEventNameFoo, nil)
	_, err = waiter.Wait()
	require.ErrorContains(t, err, "WaitForEvent can only be called once")

	waiter = newWaiter()
	waiter.WaitForEvent(emitter, testEventNameFoobar, nil)
	waiter.RejectOnEvent(emitter, testEventNameFoo, nil)
	_, err = waiter.Wait()
	require.ErrorContains(t, err, "call RejectOnEvent before WaitForEvent")
}

func TestWaiterDeadlockForErrChanCapIs1AndCallbackErr(t *testing.T) {
	timeout := 500.0
	emitter := &eventEmitter{}
	w := &waiter{
		// just receive event timeout err or callback err
		errChan: make(chan error, 1),
	}

	overCh := make(chan struct{})
	errUnReachable := errors.New("unreachable")
	err := errUnReachable

	go func() {
		_, err = w.WithTimeout(timeout).WaitForEvent(emitter, "", nil).RunAndWait(func() error {
			time.Sleep(1 * time.Second)
			close(overCh)
			//block for this err, for waiter.errChan has cache event timeout err
			return errors.New("mock timeout error")
		})
		panic("unreachable")
	}()

	<-overCh
	time.Sleep(2 * time.Second)
	require.ErrorIs(t, err, errUnReachable)
}

func TestWaiterHasNotDeadlockForErrChanCapBiggerThan1AndCallbackErr(t *testing.T) {
	timeout := 500.0
	emitter := &eventEmitter{}
	w := newWaiter().WithTimeout(timeout)
	errMockTimeout := errors.New("mock timeout error")

	var err error
	overCh := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		require.Error(t, err)
		require.NotErrorIs(t, err, errMockTimeout)
		require.ErrorIs(t, err, ErrTimeout)
		close(overCh)
	}()

	_, err = w.WaitForEvent(emitter, "", nil).RunAndWait(func() error {
		time.Sleep(1 * time.Second)
		return errMockTimeout
	})

	<-overCh
}
