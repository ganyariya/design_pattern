package golang

type Iterator interface {
	Next() interface{}
	HasNext() bool
	Reset()
}
