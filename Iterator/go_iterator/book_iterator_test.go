package go_iterator

import "testing"

func TestBookIterator(t *testing.T) {
	books := []Book{NewBook("A", 1), NewBook("B", 2)}
	bookCollection := NewBookCollection(books)
	iterator := bookCollection.createIterator()

	retrieved := []Book{}
	for iterator.HasNext() {
		element := iterator.Next()
		switch element := element.(type) {
		case Book:
			retrieved = append(retrieved, element)
		}
	}
	if len(books) != len(retrieved) {
		t.Fatal("Error")
	}

	if iterator.HasNext() {
		t.Fatal("Error")
	}

	iterator.Reset()
	if !iterator.HasNext() {
		t.Fatal("Error")
	}
}
