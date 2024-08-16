package board

import "fmt"

func GetConfig() (size int, err error) {
	fmt.Printf("Введите размер доски: ")
	_, err = fmt.Scanln(&size)
	if err != nil {
		return
	}

	return
}

func Paint(size int, black string, white string) (board string) {
	board += addtopLine(size)
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			board += "|" + paintLine(size, black, white) + "|\n"
		} else {
			board += "|" + paintLine(size, white, black) + "|\n"
		}
	}
	board += addBottomLine(size)
	return
}

func addtopLine(size int) (line string) {
	return "_" + paintLine(size, "_", "_") + "_\n"
}

func addBottomLine(size int) (line string) {
	return "⎺" + paintLine(size, "⎺", "⎺") + "⎺\n"
}

func paintLine(size int, black string, white string) (line string) {
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			line += black
		} else {
			line += white
		}
	}
	return
}
