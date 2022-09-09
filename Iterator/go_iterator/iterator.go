package go_iterator

type Iterator interface {
	Next() interface{}
	HasNext() bool
	Reset()
}
