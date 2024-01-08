package playwright

import (
	"math"
	"reflect"
	"sync"

	"golang.org/x/exp/slices"
)

type EventEmitter interface {
	Emit(name string, payload ...interface{}) bool
	ListenerCount(name string) int
	On(name string, handler interface{})
	Once(name string, handler interface{})
	RemoveListener(name string, handler interface{})
}

type (
	eventEmitter struct {
		eventsMutex sync.Mutex
		events      map[string]*eventRegister
		hasInit     bool
	}
	eventRegister struct {
		listeners []listener
	}
	listener struct {
		handler interface{}
		once    bool
	}
)

func (e *eventEmitter) Emit(name string, payload ...interface{}) (hasListener bool) {
	e.eventsMutex.Lock()
	defer e.eventsMutex.Unlock()
	e.init()

	evt, ok := e.events[name]
	if !ok {
		return
	}

	hasListener = evt.count() > 0

	evt.callHandlers(payload...)
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
	e.init()

	if _, ok := e.events[name]; !ok {
		return
	}
	e.events[name].removeHandler(handler)
}

// ListenerCount count the listeners by name, count all if name is empty
func (e *eventEmitter) ListenerCount(name string) int {
	e.eventsMutex.Lock()
	defer e.eventsMutex.Unlock()
	e.init()

	if name != "" {
		evt, ok := e.events[name]
		if !ok {
			return 0
		}
		return evt.count()
	}

	count := 0
	for key := range e.events {
		count += e.events[key].count()
	}

	return count
}

func (e *eventEmitter) addEvent(name string, handler interface{}, once bool) {
	e.eventsMutex.Lock()
	e.init()

	if _, ok := e.events[name]; !ok {
		e.events[name] = &eventRegister{
			listeners: make([]listener, 0),
		}
	}
	e.events[name].addHandler(handler, once)
	e.eventsMutex.Unlock()
}

func (e *eventEmitter) init() {
	if !e.hasInit {
		e.events = make(map[string]*eventRegister, 0)
		e.hasInit = true
	}
}

func (e *eventRegister) addHandler(handler interface{}, once bool) {
	e.listeners = append(e.listeners, listener{handler: handler, once: once})
}

func (e *eventRegister) count() int {
	return len(e.listeners)
}

func (e *eventRegister) removeHandler(handler interface{}) {
	handlerPtr := reflect.ValueOf(handler).Pointer()

	e.listeners = slices.DeleteFunc[[]listener](e.listeners, func(l listener) bool {
		return reflect.ValueOf(l.handler).Pointer() == handlerPtr
	})
}

func (e *eventRegister) callHandlers(payloads ...interface{}) {
	payloadV := make([]reflect.Value, 0)

	for _, p := range payloads {
		payloadV = append(payloadV, reflect.ValueOf(p))
	}

	handle := func(l listener) {
		handlerV := reflect.ValueOf(l.handler)
		handlerV.Call(payloadV[:int(math.Min(float64(handlerV.Type().NumIn()), float64(len(payloadV))))])
	}

	for _, l := range e.listeners {
		if l.once {
			defer e.removeHandler(l.handler)
		}
		handle(l)
	}
}
