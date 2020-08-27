package playwright

import (
	"reflect"
	"sync"
)

type (
	incomingEvent struct {
		name    string
		payload []interface{}
	}
	eventHandler  func(...interface{})
	eventRegister struct {
		once []eventHandler
		on   []eventHandler
	}
	EventEmitter struct {
		sync.Mutex
		queue  chan incomingEvent
		events map[string]*eventRegister
	}
)

func (e *EventEmitter) Emit(name string, payload ...interface{}) {
	e.queue <- incomingEvent{name, payload}
}

func (e *EventEmitter) Once(name string, handler eventHandler) {
	e.addEvent(name, handler, true)
}

func (e *EventEmitter) On(name string, handler eventHandler) {
	e.addEvent(name, handler, false)
}

func (e *EventEmitter) RemoveListener(name string, handler eventHandler) {
	e.Lock()
	defer e.Unlock()
	if _, ok := e.events[name]; !ok {
		return
	}
	handlerPtr := reflect.ValueOf(handler).Pointer()
	onHandlers := []eventHandler{}
	for idx := range e.events[name].on {
		eventPtr := reflect.ValueOf(e.events[name].on[idx]).Pointer()
		if eventPtr != handlerPtr {
			onHandlers = append(onHandlers, e.events[name].on[idx])
		}
	}

	onceHandlers := []eventHandler{}
	for idx := range e.events[name].once {
		eventPtr := reflect.ValueOf(e.events[name].once[idx]).Pointer()
		if eventPtr != handlerPtr {
			onceHandlers = append(onceHandlers, e.events[name].once[idx])
		}
	}

	e.events[name].on = onHandlers
	e.events[name].once = onceHandlers
}

func (e *EventEmitter) ListenerCount(name string) int {
	count := 0
	e.Lock()
	for key := range e.events {
		count += len(e.events[key].on) + len(e.events[key].once)
	}
	e.Unlock()
	return count
}

func (e *EventEmitter) addEvent(name string, handler eventHandler, once bool) {
	e.Lock()
	if _, ok := e.events[name]; !ok {
		e.events[name] = &eventRegister{
			on:   make([]eventHandler, 0),
			once: make([]eventHandler, 0),
		}
	}
	if once {
		e.events[name].once = append(e.events[name].once, handler)
	} else {
		e.events[name].on = append(e.events[name].on, handler)
	}
	e.Unlock()
}

func (e *EventEmitter) initEventEmitter() {
	e.events = make(map[string]*eventRegister)
	e.queue = make(chan incomingEvent)
	go e.startEventQueue()
}

func (e *EventEmitter) startEventQueue() {
	for {
		payload, more := <-e.queue
		if !more {
			break
		}
		e.Lock()
		if _, ok := e.events[payload.name]; !ok {
			e.Unlock()
			continue
		}

		for i := 0; i < len(e.events[payload.name].on); i++ {
			e.events[payload.name].on[i](payload.payload...)
		}
		for i := 0; i < len(e.events[payload.name].once); i++ {
			e.events[payload.name].once[i](payload.payload...)
		}
		e.events[payload.name].once = make([]eventHandler, 0)
		e.Unlock()
	}
}

func (e *EventEmitter) stopEventEmitter() {
	close(e.queue)
}
