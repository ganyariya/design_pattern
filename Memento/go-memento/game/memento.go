package game

/*
セーブデータ
*/
type Memento struct {
	counter int
}

func newMemento(counter int) *Memento {
	return &Memento{counter: counter}
}
