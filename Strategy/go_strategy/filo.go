package go_strategy

type FILO[T any] struct{}

func (f *FILO[T]) apply(c *Cache[T]) T {
	if len(c.data) == 0 {
		return *new(T)
	}
	n := len(c.data)
	x := c.data[n-1]
	c.data = c.data[0 : n-1]
	return x
}
