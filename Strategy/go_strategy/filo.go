package go_strategy

type FILO struct{}

func (f *FILO) apply(c *Cache) int {
	if len(c.data) == 0 {
		return -1
	}
	n := len(c.data)
	x := c.data[n-1]
	c.data = c.data[0 : n-1]
	return x
}
