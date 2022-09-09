package go_iterator

type BookIterator struct {
	bookCollection *BookCollection
	index          int
}

func NewBookIterator(bookCollection *BookCollection) *BookIterator {
	return &BookIterator{bookCollection: bookCollection}
}

func (bi *BookIterator) Next() interface{} {
	res := bi.bookCollection.books[bi.index]
	bi.index++
	return res
}

func (bi *BookIterator) HasNext() bool {
	return bi.index < len(bi.bookCollection.books)
}

func (bi *BookIterator) Reset() {
	bi.index = 0
}
