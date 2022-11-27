package visitor

type Entry interface {
	GetSize() int
	GetName() string
	Acceptor
}

type Acceptor interface {
	Accept(v VisitorInterface) any
}
