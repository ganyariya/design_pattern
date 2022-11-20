package go_strategy

type Strategy interface {
	apply(c *Cache) int
}
