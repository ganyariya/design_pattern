package gomediator

type EventDispatcherInterface interface {
	Attach(listener EventListenerInterface, event string)
	Dispatch(event string, emitter EventListenerInterface, data any)
}

type EventDispatcher struct {
	Listeners map[string][]EventListenerInterface
}

func NewEventDispatcher() EventDispatcherInterface {
	return &EventDispatcher{Listeners: map[string][]EventListenerInterface{}}
}

func (ed *EventDispatcher) initObserverSlice(event string) {
	if _, ok := ed.Listeners[event]; !ok {
		ed.Listeners[event] = []EventListenerInterface{}
	}
}
func (ed *EventDispatcher) Attach(observer EventListenerInterface, event string) {
	ed.initObserverSlice(event)
	ed.Listeners[event] = append(ed.Listeners[event], observer)
}
func (ed *EventDispatcher) Dispatch(event string, emitter EventListenerInterface, data any) {
	ed.initObserverSlice(event)
	for _, listener := range ed.Listeners[event] {
		listener.Receive(event, emitter, data)
	}
}
