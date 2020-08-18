package playwright

import "sync"

type eventHandler = func(payload ...interface{})

type eventRegister struct {
	sync.RWMutex
	once []eventHandler
	on   []eventHandler
}

type EventEmitter struct {
	events map[string]*eventRegister
}

func (e *EventEmitter) initEventEmitter() {
	e.events = make(map[string]*eventRegister)
}

func (e *EventEmitter) Emit(name string, payload ...interface{}) {
	if _, ok := e.events[name]; ok {
		e.events[name].RLock()
		for i := 0; i < len(e.events[name].on); i++ {
			go e.events[name].on[i](payload...)
		}
		for i := 0; i < len(e.events[name].once); i++ {
			go e.events[name].once[i](payload...)
		}
		e.events[name].once = make([]eventHandler, 0)
		e.events[name].RUnlock()
	}
}

func (e *EventEmitter) Once(name string, handler eventHandler) {
	e.addEvent(name, handler, true)
}

func (e *EventEmitter) On(name string, handler eventHandler) {
	e.addEvent(name, handler, false)
}

func (e *EventEmitter) RemoveListener(name string, handler eventHandler) {
	if _, ok := e.events[name]; !ok {
		return
	}
	onHandlers := []eventHandler{}
	for idx := range e.events[name].on {
		if &e.events[name].on[idx] != &handler {
			onHandlers = append(onHandlers, e.events[name].on[idx])
		}
	}
	e.events[name].Lock()
	e.events[name].on = onHandlers
	e.events[name].Unlock()

	onceHandlers := []eventHandler{}
	for idx := range e.events[name].once {
		if &e.events[name].once[idx] != &handler {
			onceHandlers = append(onceHandlers, e.events[name].once[idx])
		}
	}
	e.events[name].Lock()
	e.events[name].once = onceHandlers
	e.events[name].Unlock()
}

func (e *EventEmitter) addEvent(name string, handler eventHandler, once bool) {
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
