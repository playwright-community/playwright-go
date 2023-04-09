package playwright

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

type waiter struct {
	mu        sync.Mutex
	err       error
	timeout   float64
	fulfill   int32
	listeners []eventListener
	errChan   chan error
}

type eventListener struct {
	emitter EventEmitter
	event   string
	handler interface{}
}

func (w *waiter) reject(err error) {
	atomic.StoreInt32(&w.fulfill, 1)
	w.errChan <- err
}

func (w *waiter) Err() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.err
}

func (w *waiter) RejectOnEvent(emitter EventEmitter, event string, err error, predicate ...interface{}) {
	w.mu.Lock()
	defer w.mu.Unlock()
	handler := func(ev ...interface{}) {
		if atomic.LoadInt32(&w.fulfill) == 1 {
			return
		}
		if len(predicate) == 0 {
			w.reject(err)
		} else if len(predicate) == 1 {
			result := reflect.ValueOf(predicate[0]).Call([]reflect.Value{reflect.ValueOf(ev[0])})
			if result[0].Bool() {
				w.reject(err)
			}
		}
	}
	emitter.On(event, handler)
	w.listeners = append(w.listeners, eventListener{
		emitter: emitter,
		event:   event,
		handler: handler,
	})
}

func (w *waiter) RejectOnTimeout(timeout float64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.timeout = timeout
}

func (w *waiter) WaitForEvent(emitter EventEmitter, event string, predicate ...interface{}) <-chan interface{} {
	w.mu.Lock()
	defer w.mu.Unlock()
	evChan := make(chan interface{}, 1)
	handler := w.createHandler(evChan, predicate...)
	ctx, cancel := context.WithCancel(context.Background())
	if w.timeout != 0 {
		timeout := w.timeout
		go func() {
			select {
			case <-time.After(time.Duration(timeout) * time.Millisecond):
				err := &TimeoutError{
					Name:    "TimeError",
					Message: fmt.Sprintf("Timeout %.2fms exceeded.", timeout),
				}
				w.reject(err)
				return
			case <-ctx.Done():
				return
			}
		}()
	}
	go func() {
		err := <-w.errChan
		cancel()
		w.mu.Lock()
		defer w.mu.Unlock()
		for _, l := range w.listeners {
			l.emitter.RemoveListener(l.event, l.handler)
		}
		close(evChan)
		w.err = err
	}()

	emitter.On(event, handler)
	w.listeners = append(w.listeners, eventListener{
		emitter: emitter,
		event:   event,
		handler: handler,
	})

	return evChan
}

func (w *waiter) createHandler(evChan chan<- interface{}, predicate ...interface{}) func(...interface{}) {
	return func(ev ...interface{}) {
		if atomic.LoadInt32(&w.fulfill) == 1 {
			return
		}
		if len(predicate) == 0 {
			if len(ev) == 1 {
				evChan <- ev[0]
			} else {
				evChan <- nil
			}
			w.reject(nil)
		} else if len(predicate) == 1 {
			result := reflect.ValueOf(predicate[0]).Call([]reflect.Value{reflect.ValueOf(ev[0])})
			if result[0].Bool() {
				evChan <- ev[0]
				w.reject(nil)
			}
		}
	}
}

func newWaiter() *waiter {
	w := &waiter{
		errChan: make(chan error, 1),
	}
	return w
}
