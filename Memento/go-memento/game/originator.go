package game

/*
Game 内容
*/
type Originator struct {
	counter int
}

func NewOriginator(counter int) *Originator {
	return &Originator{counter: counter}
}
func (o *Originator) Play() {
	o.counter++
}
func (o *Originator) GetCounter() int {
	return o.counter
}
func (o *Originator) CreateMemento() *Memento {
	return newMemento(o.counter)
}
func (o *Originator) RestoreMemento(m *Memento) {
	o.counter = m.counter
}
