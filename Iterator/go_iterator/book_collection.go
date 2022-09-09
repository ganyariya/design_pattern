package go_iterator

type BookCollection struct {
	books []Book
}

func NewBookCollection(books []Book) *BookCollection {
	return &BookCollection{books: books}
}

func (bc *BookCollection) createIterator() Iterator {
	return NewBookIterator(bc)
}

func (bc *BookCollection) Add(book Book) {
	bc.books = append(bc.books, book)
}
