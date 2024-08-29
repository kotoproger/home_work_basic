package main

import (
	"fmt"

	"github.com/kotoproger/hw04_struct_comparator/book"
	"github.com/kotoproger/hw04_struct_comparator/comparator"
)

func main() {
	book1 := book.NewBook(123, "title1", "author1", 2000, 34, 9.5)
	book2 := book.NewBook(124, "title2", "author2", 2000, 35, 9.3)

	comparator := comparator.NewBookComparator([]string{comparator.CompareYear})

	fmt.Println("book 1:", book1)
	fmt.Println("book 2:", book2)
	fmt.Print("compare result: ", comparator.Compare(book1, book2))
}
