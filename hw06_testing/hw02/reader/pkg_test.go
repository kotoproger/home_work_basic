package reader

import (
	"fmt"
	"os"
	"testing"

	types "github.com/kotoproger/home_work_basic/hw06_testing/hw02/types"
	"github.com/stretchr/testify/assert"
)

func TestReadJSON(t *testing.T) {
	testCases := []struct {
		name        string
		jsonContent string
		isError     bool
		data        []types.Employee
	}{
		{"file not exists", "", true, nil},
		{"file content is invalid", "some not json content", true, nil},
		{
			"file is valid",
			"[{\"userId\":10,\"age\":25,\"name\":\"Rob\",\"departmentId\":3}," +
				"{\"userId\":11,\"age\":30,\"name\":\"George\",\"departmentId\":2}]",
			false,
			[]types.Employee{
				{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
				{UserID: 11, Age: 30, Name: "George", DepartmentID: 2},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var fileName string
			var err error
			if tc.jsonContent != "" {
				fileName, err = createTempFile(tc.jsonContent)
				defer os.Remove(fileName)
				assert.Equal(t, nil, err, "erro on creating temporary file")
			} else {
				fileName = "some random name"
			}

			list, err := ReadJSON(fileName)

			assert.Equal(t, tc.data, list)
			assert.Equal(t, tc.isError, err != nil)
		})
	}
}

func createTempFile(content string) (string, error) {
	f, err := os.CreateTemp("", "example")
	if err != nil {
		return "", fmt.Errorf("can not create temporary file: %w", err)
	}

	_, err = f.Write([]byte(content))
	if err != nil {
		return "", fmt.Errorf("can not write to temporary file: %w", err)
	}

	err = f.Close()
	if err != nil {
		return "", fmt.Errorf("can not close temporary file: %w", err)
	}

	return f.Name(), nil
}
