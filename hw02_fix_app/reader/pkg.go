package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	types "github.com/kotoproger/home_work_basic/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)

		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)

		return nil, err
	}

	var data []types.Employee

	unmarshalErr := json.Unmarshal(bytes, &data)
	if unmarshalErr != nil {
		fmt.Printf("Error: %v", unmarshalErr)

		return nil, unmarshalErr
	}

	return data, nil
}
