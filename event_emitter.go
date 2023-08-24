package playwright

import (
	"math"
	"reflect"
	"sync"
)

type EventEmitter interface {
	Emit(name string, payload ...interface{}) bool
	ListenerCount(name string) int
	On(name string, handler interface{})
	Once(name string, handler interface{})
	RemoveListener(name string, handler interface{})
}

type (
	eventRegister struct {
		once []interface{}
		on   []interface{}
	}
	eventEmitter struct {
		eventsMutex sync.Mutex
		events      map[string]*eventRegister
	}
)

func (e *eventEmitter) Emit(name string, payload ...interface{}) (handled bool) {
	e.eventsMutex.Lock()
	defer e.eventsMutex.Unlock()
	if _, ok := e.events[name]; !ok {
		return
	}

	if len(e.events[name].once) > 0 || len(e.events[name].on) > 0 {
		handled = true
	}

	payloadV := make([]reflect.Value, 0)

	for _, p := range payload {
		payloadV = append(payloadV, reflect.ValueOf(p))
	}

	callHandlers := func(handlers []interface{}) {
		for _, handler := range handlers {
			handlerV := reflect.ValueOf(handler)
			handlerV.Call(payloadV[:int(math.Min(float64(handlerV.Type().NumIn()), float64(len(payloadV))))])
		}
	}

	callHandlers(e.events[name].on)
	callHandlers(e.events[name].once)

	e.events[name].once = make([]interface{}, 0)
	return
}

func (e *eventEmitter) Once(name string, handler interface{}) {
	e.addEvent(name, handler, true)
}

func (e *eventEmitter) On(name string, handler interface{}) {
	e.addEvent(name, handler, false)
}

func (e *eventEmitter) RemoveListener(name string, handler interface{}) {
	e.eventsMutex.Lock()
	defer e.eventsMutex.Unlock()
	if _, ok := e.events[name]; !ok {
		return
	}
	handlerPtr := reflect.ValueOf(handler).Pointer()

	onHandlers := []interface{}{}
	for idx := range e.events[name].on {
		eventPtr := reflect.ValueOf(e.events[name].on[idx]).Pointer()
		if eventPtr != handlerPtr {
			onHandlers = append(onHandlers, e.events[name].on[idx])
		}
	}
	e.events[name].on = onHandlers

	onceHandlers := []interface{}{}
	for idx := range e.events[name].once {
		eventPtr := reflect.ValueOf(e.events[name].once[idx]).Pointer()
		if eventPtr != handlerPtr {
			onceHandlers = append(onceHandlers, e.events[name].once[idx])
		}
	}

	e.events[name].once = onceHandlers
}

// ListenerCount count the listeners by name, count all if name is empty
func (e *eventEmitter) ListenerCount(name string) int {
	count := 0
	e.eventsMutex.Lock()
	for key := range e.events {
		if name == "" || name == key {
			count += len(e.events[key].on) + len(e.events[key].once)
		}
	}
	e.eventsMutex.Unlock()
	return count
}

func (e *eventEmitter) addEvent(name string, handler interface{}, once bool) {
	e.eventsMutex.Lock()
	if _, ok := e.events[name]; !ok {
		e.events[name] = &eventRegister{
			on:   make([]interface{}, 0),
			once: make([]interface{}, 0),
		}
	}
	if once {
		e.events[name].once = append(e.events[name].once, handler)
	} else {
		e.events[name].on = append(e.events[name].on, handler)
	}
	e.eventsMutex.Unlock()
}

func (e *eventEmitter) initEventEmitter() {
	e.events = make(map[string]*eventRegister)
}
