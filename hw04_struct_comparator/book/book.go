package book

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

func NewBook(id int,
	title string,
	author string,
	year int,
	size int,
	rate float64) Book {
	return Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func (book Book) ID() int {
	return book.id
}

func (book Book) Title() string {
	return book.title
}

func (book Book) Author() string {
	return book.author
}

func (book Book) Year() int {
	return book.year
}

func (book Book) Size() int {
	return book.size
}

func (book Book) Rate() float64 {
	return book.rate
}

func (book *Book) SetTitle(newTitle string) {
	book.title = newTitle
}

func (book *Book) SetAuthor(newAuthor string) {
	book.author = newAuthor
}

func (book *Book) SetYear(newYear int) {
	book.year = newYear
}

func (book *Book) SetSize(newSize int) {
	book.size = newSize
}

func (book *Book) SetRate(newRate float64) {
	book.rate = newRate
}
