package go_strategy

// Context
type Cache struct {
	data     []int
	strategy Strategy
}

func (c *Cache) SetStrategy(s Strategy) {
	c.strategy = s
}
func (c *Cache) Push(x int) {
	c.data = append(c.data, x)
}
func (c *Cache) GetElement() int {
	return c.strategy.apply(c)
}

func NewCache() *Cache {
	return &Cache{data: []int{}, strategy: nil}
}
