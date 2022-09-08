package golang

type Book struct {
	Name string
	Page int
}

func NewBook(name string, page int) Book {
	return Book{name, page}
}
