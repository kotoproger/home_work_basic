package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	types "github.com/kotoproger/home_work_basic/hw06_testing/hw02/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open json file: %w", err)
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read json file: %w", err)
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json string: %w", err)
	}

	return data, nil
}
