package go_strategy

type Strategy[T any] interface {
	apply(c *Cache[T]) T
}
