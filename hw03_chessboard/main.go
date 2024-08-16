package main

import (
	"fmt"
	"log"

	"github.com/kotoproger/home_work_basic/hw03_chessboard/board"
)

func main() {
	size, err := board.GetConfig()
	if err != nil {
		log.Fatal(err)

		return
	}

	board := board.Paint(size, "#", " ")

	fmt.Println(board)
}
