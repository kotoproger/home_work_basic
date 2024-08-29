package comparator

import (
	"github.com/kotoproger/home_work_basic/hw04_struct_comparator/book"
)

type CompareType string

const (
	CompareYear CompareType = "year"
	CompareSize CompareType = "size"
	CompareRate CompareType = "rate"
)

type BookComparator struct {
	compareType CompareType
}

func NewBookComparator(compareType CompareType) *BookComparator {
	return &BookComparator{
		compareType: compareType,
	}
}

func (comp BookComparator) Compare(book1 *book.Book, book2 *book.Book) bool {
	switch comp.compareType {
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

	return false
}
