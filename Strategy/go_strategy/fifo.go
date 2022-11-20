package go_strategy

type FIFO struct{}

func (f *FIFO) apply(c *Cache) int {
	if len(c.data) == 0 {
		return -1
	}
	x := c.data[0]
	c.data = c.data[1:]
	return x
}
