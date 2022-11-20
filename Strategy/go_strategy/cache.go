package go_strategy

// Context
type Cache[T any] struct {
	data     []T
	strategy Strategy[T]
}

func (c *Cache[T]) SetStrategy(s Strategy[T]) {
	c.strategy = s
}
func (c *Cache[T]) Push(x T) {
	c.data = append(c.data, x)
}
func (c *Cache[T]) GetElement() T {
	return c.strategy.apply(c)
}

func NewCache[T any]() *Cache[T] {
	return &Cache[T]{data: []T{}, strategy: nil}
}
