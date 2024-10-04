package book

import (
	"encoding/json"
	"fmt"
)

type jsonBook struct {
	ID     uint32  `json:"idid"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   uint32  `json:"year"`
	Size   uint32  `json:"size"`
	Rate   float32 `json:"rate"`
}

type jsonBooklist struct {
	Books []jsonBook `json:"bOOks"`
}

func (jb *jsonBook) book() Book {
	return Book{
		ID:     jb.ID,
		Title:  jb.Title,
		Author: jb.Author,
		Year:   jb.Year,
		Size:   jb.Size,
		Rate:   jb.Rate,
	}
}

func (jb *jsonBook) toBook(b *Book) {
	b.ID = jb.ID
	b.Title = jb.Title
	b.Author = jb.Author
	b.Year = jb.Year
	b.Size = jb.Size
	b.Rate = jb.Rate
}

func (b *Book) jsonBook() (jb jsonBook) {
	jb.ID = b.ID
	jb.Title = b.Title
	jb.Author = b.Author
	jb.Year = b.Year
	jb.Size = b.Size
	jb.Rate = b.Rate
	return
}

func (b Book) MarshalJSON() ([]byte, error) {
	fmt.Println("book marshaler")
	return json.Marshal(b.jsonBook())
}

func (b *Book) UnmarshalJSON(data []byte) error {
	jBook := jsonBook{}
	err := json.Unmarshal(
		data,
		&jBook,
	)
	if err != nil {
		return err
	}

	jBook.toBook(b)

	return nil
}

func (bl Booklist) MarshalJSON() ([]byte, error) {
	jbl := jsonBooklist{Books: make([]jsonBook, len(bl.Books))}
	for index, book := range bl.Books {
		jbl.Books[index] = book.jsonBook()
	}

	return json.Marshal(jbl)
}

func (bl *Booklist) UnmarshalJSON(data []byte) error {
	jbl := jsonBooklist{}
	err := json.Unmarshal(data, &jbl)
	if err != nil {
		return err
	}
	bl.Books = make([]*Book, len(jbl.Books))
	for index, jbook := range jbl.Books {
		book := jbook.book()
		bl.Books[index] = &book
	}

	return nil
}
