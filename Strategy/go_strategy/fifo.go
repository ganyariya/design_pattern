package go_strategy

type FIFO[T any] struct{}

func (f *FIFO[T]) apply(c *Cache[T]) T {
	if len(c.data) == 0 {
		return *new(T)
	}
	x := c.data[0]
	c.data = c.data[1:]
	return x
}
