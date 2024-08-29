package comparator

import (
	"github.com/fixme_my_friend/hw04_struct_comparator/book"
)

const (
	COMPARE_YEAR = "year"
	COMPARE_SIZE = "size"
	COMPARE_RATE = "rate"
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
		case COMPARE_RATE:
			if book1.Rate() > book2.Rate() {
				return true
			}
		case COMPARE_SIZE:
			if book1.Size() > book2.Size() {
				return true
			}
		case COMPARE_YEAR:
			if book1.Year() > book2.Year() {
				return true
			}
		}
	}

	return false
}
