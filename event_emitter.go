package playwright

import "sync"

type (
	eventName     string
	incomingEvent struct {
		name    eventName
		payload []interface{}
	}
	eventHandler  func(...interface{})
	eventRegister struct {
		sync.Mutex
		once []eventHandler
		on   []eventHandler
	}
	EventEmitter struct {
		queue  chan incomingEvent
		events map[eventName]*eventRegister
	}
)

func (e *EventEmitter) Emit(name eventName, payload ...interface{}) {
	if _, ok := e.events[name]; ok {
		e.queue <- incomingEvent{name, payload}
	}
}

func (e *EventEmitter) Once(name eventName, handler eventHandler) {
	e.addEvent(name, handler, true)
}

func (e *EventEmitter) On(name eventName, handler eventHandler) {
	e.addEvent(name, handler, false)
}

func (e *EventEmitter) RemoveListener(name eventName, handler eventHandler) {
	if _, ok := e.events[name]; !ok {
		return
	}
	onHandlers := []eventHandler{}
	for idx := range e.events[name].on {
		if &e.events[name].on[idx] != &handler {
			onHandlers = append(onHandlers, e.events[name].on[idx])
		}
	}

	onceHandlers := []eventHandler{}
	for idx := range e.events[name].once {
		if &e.events[name].once[idx] != &handler {
			onceHandlers = append(onceHandlers, e.events[name].once[idx])
		}
	}

	e.events[name].Lock()
	e.events[name].on = onHandlers
	e.events[name].once = onceHandlers
	e.events[name].Unlock()
}

func (e *EventEmitter) ListenerCount(name string) int {
	count := 0
	for key := range e.events {
		e.events[key].Lock()
		count += len(e.events[key].on) + len(e.events[key].once)
		e.events[key].Unlock()
	}
	return count
}

func (e *EventEmitter) addEvent(name eventName, handler eventHandler, once bool) {
	if _, ok := e.events[name]; !ok {
		e.events[name] = &eventRegister{
			on:   make([]eventHandler, 0),
			once: make([]eventHandler, 0),
		}
	}
	e.events[name].Lock()
	if once {
		e.events[name].once = append(e.events[name].once, handler)
	} else {
		e.events[name].on = append(e.events[name].on, handler)
	}
	e.events[name].Unlock()
}

func (e *EventEmitter) initEventEmitter() {
	e.events = make(map[eventName]*eventRegister)
	e.queue = make(chan incomingEvent)
	go e.startEventQueue()
}

func (e *EventEmitter) startEventQueue() {
	for {
		payload, more := <-e.queue
		if !more {
			break
		}
		if _, ok := e.events[payload.name]; !ok {
			continue
		}

		e.events[payload.name].Lock()
		for i := 0; i < len(e.events[payload.name].on); i++ {
			e.events[payload.name].on[i](payload.payload...)
		}
		for i := 0; i < len(e.events[payload.name].once); i++ {
			e.events[payload.name].once[i](payload.payload...)
		}
		e.events[payload.name].once = make([]eventHandler, 0)
		e.events[payload.name].Unlock()
	}
}

func (e *EventEmitter) stopEventEmitter() {
	close(e.queue)
}
