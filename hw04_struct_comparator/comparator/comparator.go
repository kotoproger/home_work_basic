package comparator

import (
	"github.com/kotoproger/hw04_struct_comparator/book"
)

const (
	CompareYear = "year"
	CompareSize = "size"
	CompareRate = "rate"
)

type BookComparator struct {
	compareTypes []string
}

func NewBookComparator(types []string) BookComparator {
	return BookComparator{
		compareTypes: types,
	}
}

func (comp BookComparator) Compare(book1 *book.Book, book2 *book.Book) bool {
	for _, field := range comp.compareTypes {
		switch field {
		case CompareRate:
			if book1.Rate() > book2.Rate() {
				return true
			}
		case CompareSize:
			if book1.Size() > book2.Size() {
				return true
			}
		case CompareYear:
			if book1.Year() > book2.Year() {
				return true
			}
		}
	}

	return false
}
