package composite

type Entry interface {
	GetSize() int
	GetName() string
	GetList(prefix string) string
}
