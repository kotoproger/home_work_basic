package main

import (
	"encoding/json"
	"fmt"

	"github.com/kotoproger/home_work_basic/hw09_serialize/book"
)

func main() {
	b := book.Book{}
	fmt.Println(json.Unmarshal([]byte("{}"), &b))
	fmt.Print(b)
}
