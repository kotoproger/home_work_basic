package configapp

import (
	"os"
	"strings"
)

func parseEnv(values map[string]*string, params map[string]ConfigParam) map[string]*string {
	for _, param := range params {
		value, ok := os.LookupEnv("APP_" + strings.ToUpper(param.Name))
		if ok {
			if param.IsBool {
				value = "true"
			}
			values[param.Name] = &value
		} else if param.IsBool {
			value = "false"
			values[param.Name] = &value
		}
	}

	return values
}
