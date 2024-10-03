package book

import (
	"encoding/json"
	"strings"
)

type jsonBook struct {
	ID     uint32  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   uint32  `json:"year"`
	Size   uint32  `json:"size"`
	Rate   float32 `json:"rate"`
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

func (jb *jsonBook) fromBook(b *Book) {
	jb.ID = b.ID
	jb.Title = b.Title
	jb.Author = b.Author
	jb.Year = b.Year
	jb.Size = b.Size
	jb.Rate = b.Rate
}

func (b *Book) MarshalJSON() ([]byte, error) {
	jBook := jsonBook{}
	jBook.fromBook(b)
	return json.Marshal(jBook)
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

func MarshalBookSlice(books []Book) ([]byte, error) {
	var b strings.Builder

	_, startErr := b.WriteString("[")
	if startErr != nil {
		return nil, startErr
	}

	first := true
	for _, el := range books {
		if first {
			first = false
		} else {
			_, err := b.WriteString(",")
			if err != nil {
				return nil, err
			}
		}
		jsonString, mErr := el.MarshalJSON()
		if mErr != nil {
			return nil, mErr
		}
		b.WriteString(string(jsonString))
	}

	_, endErr := b.WriteString("]")
	if endErr != nil {
		return nil, endErr
	}

	return []byte(b.String()), nil
}

func UnmarshalBookSliceJSON(data []byte) ([]Book, error) {
	tempStruct := []jsonBook{}
	err := json.Unmarshal(
		data,
		&tempStruct,
	)
	if err != nil {
		return nil, err
	}
	result := make([]Book, len(tempStruct))
	for index, ts := range tempStruct {
		result[index] = ts.book()
	}

	return result, nil
}
