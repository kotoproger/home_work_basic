package book

import (
	"encoding/json"
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
