package main

import (
	"fmt"
	"log"
)

func main() {
	size, err := getBoardConfig()
	if err != nil {
		log.Fatal(err)

		return
	}

	board := paintBoard(size, "#", " ")

	fmt.Println(board)
}

func getBoardConfig() (size int, err error) {
	fmt.Printf("Введите размер доски: ")
	_, err = fmt.Scanln(&size)
	if err != nil {
		return
	}

	return
}

func paintBoard(size int, black string, white string) (board string) {
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			board += paintLine(size, black, white)
		} else {
			board += paintLine(size, white, black)
		}
	}

	return
}

func paintLine(size int, black string, white string) (line string) {
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			line += black
		} else {
			line += white
		}
	}

	line += "\n"
	return
}
