package gomediator

type EventListenerInterface interface {
	Receive(event string, emitter EventListenerInterface, data any)
}

type EventListenerForm struct {
	dispatcher      EventDispatcherInterface
	ActualData      string
	NotifiedMessage string
}

func NewEventListenerForm(dispatcher EventDispatcherInterface, event string) *EventListenerForm {
	elf := &EventListenerForm{dispatcher: dispatcher}
	dispatcher.Attach(elf, event)
	return elf
}

func (elf *EventListenerForm) Receive(event string, emitter EventListenerInterface, data any) {
	elf.NotifiedMessage = data.(string)
}

type EventListenerButton struct {
	dispatcher EventDispatcherInterface
}

func NewEventListenerButton(dispatcher EventDispatcherInterface, event string) *EventListenerButton {
	elb := &EventListenerButton{
		dispatcher: dispatcher,
	}
	dispatcher.Attach(elb, event)
	return elb
}

func (elb *EventListenerButton) Receive(event string, emitter EventListenerInterface, data any) {}
func (elb *EventListenerButton) Click(event string) {
	elb.dispatcher.Dispatch(event, elb, "button submitted")
}
