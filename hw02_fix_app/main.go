package main

import (
	"fmt"

	"github.com/kotoproger/home_work_basic/hw02_fix_app/printer"
	"github.com/kotoproger/home_work_basic/hw02_fix_app/reader"
	"github.com/kotoproger/home_work_basic/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	_, scanErr := fmt.Scanln(&path)
	if scanErr != nil {
		fmt.Printf("Error: %v", scanErr)
		return
	}

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	printer.PrintStaff(staff)
}
