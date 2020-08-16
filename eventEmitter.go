package playwright

type event struct {
	function func(payload ...interface{})
	once     bool
}

type eventRegister struct {
	events []*event
}

type EventEmitter struct {
	events map[string]*eventRegister
}

func (e *EventEmitter) initEventEmitter() {
	e.events = make(map[string]*eventRegister)
}

func (e *EventEmitter) Emit(name string, payload ...interface{}) {
	if _, ok := e.events[name]; ok {
		for i := 0; i < len(e.events[name].events); i++ {
			go e.events[name].events[i].function(payload...)
			if e.events[name].events[i].once == true {
				e.events[name].events[i] = nil
			}
			e.events[name].Cleanup()
		}
	}
}

func (e *EventEmitter) Once(name string, handler func(payload ...interface{})) {
	e.addEvent(name, handler, true)
}

func (e *EventEmitter) On(name string, handler func(payload ...interface{})) {
	e.addEvent(name, handler, false)
}
func (e *EventEmitter) addEvent(name string, handler func(payload ...interface{}), once bool) {
	if _, ok := e.events[name]; !ok {
		e.events[name] = &eventRegister{
			events: make([]*event, 0),
		}
	}
	e.events[name].events = append(e.events[name].events, &event{
		function: handler,
		once:     once,
	})
}

func (e *eventRegister) Cleanup() {
	events := make([]*event, 0)
	for i := 0; i < len(e.events); i++ {
		if e.events[i] != nil {
			events = append(events, e.events[i])
		}
	}
	e.events = events
}
