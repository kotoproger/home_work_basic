package book

import (
	"encoding/json"
	"strings"
)

func (b *Book) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			ID     uint32  `json:"id"`
			Title  string  `json:"title"`
			Author string  `json:"author"`
			Year   uint32  `json:"year"`
			Size   uint32  `json:"size"`
			Rate   float32 `json:"rate"`
		}{
			b.ID,
			b.Title,
			b.Author,
			b.Year,
			b.Size,
			b.Rate,
		},
	)
}

func (b *Book) UnmarshalJSON(data []byte) error {
	tempStruct := struct {
		ID     uint32  `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   uint32  `json:"year"`
		Size   uint32  `json:"size"`
		Rate   float32 `json:"rate"`
	}{}
	err := json.Unmarshal(
		data,
		&tempStruct,
	)
	if err != nil {
		return err
	}

	b.ID = tempStruct.ID
	b.Title = tempStruct.Title
	b.Author = tempStruct.Author
	b.Year = tempStruct.Year
	b.Size = tempStruct.Size
	b.Rate = tempStruct.Rate

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
	tempStruct := []struct {
		ID     uint32  `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   uint32  `json:"year"`
		Size   uint32  `json:"size"`
		Rate   float32 `json:"rate"`
	}{}
	err := json.Unmarshal(
		data,
		&tempStruct,
	)
	if err != nil {
		return nil, err
	}
	result := make([]Book, len(tempStruct))
	for index, ts := range tempStruct {
		result[index] = Book{
			ID:     ts.ID,
			Title:  ts.Title,
			Author: ts.Author,
			Year:   ts.Year,
			Size:   ts.Size,
			Rate:   ts.Rate,
		}
	}

	return result, nil
}
