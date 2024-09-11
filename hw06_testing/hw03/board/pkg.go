package board

import (
	"fmt"
	"strings"
)

func GetConfig() (size int, err error) {
	fmt.Printf("Введите размер доски: ")
	_, err = fmt.Scanln(&size)
	if err != nil {
		return
	}

	return
}

func Paint(size int, black string, white string) (board string) {
	var b strings.Builder
	b.WriteString(addtopLine(size))
	for i := 0; i < size; i++ {
		b.WriteString("|")
		if i%2 == 0 {
			b.WriteString(paintLine(size, black, white))
		} else {
			b.WriteString(paintLine(size, white, black))
		}
		b.WriteString("|\n")
	}
	b.WriteString(addBottomLine(size))

	return b.String()
}

func addtopLine(size int) (line string) {
	var b strings.Builder
	b.WriteString("_")
	b.WriteString(paintLine(size, "_", "_"))
	b.WriteString("_\n")

	return b.String()
}

func addBottomLine(size int) (line string) {
	var b strings.Builder
	b.WriteString("⎺")
	b.WriteString(paintLine(size, "⎺", "⎺"))
	b.WriteString("⎺\n")

	return b.String()
}

func paintLine(size int, black string, white string) (line string) {
	var b strings.Builder
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			b.WriteString(black)
		} else {
			b.WriteString(white)
		}
	}

	return b.String()
}
