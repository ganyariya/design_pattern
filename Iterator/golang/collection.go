package golang

type Collection interface {
	createIterator() Iterator
}
