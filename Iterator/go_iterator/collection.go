package go_iterator

type Collection interface {
	createIterator() Iterator
}
