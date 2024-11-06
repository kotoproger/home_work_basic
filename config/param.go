package configapp

import (
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

func parseParams(values map[string]*string, params map[string]ConfigParam) map[string]*string {
	boolMapping := make(map[string]*bool)
	for _, param := range params {
		flag := pflag.Lookup(strings.ToLower(param.Name))
		if flag == nil {
			if param.IsBool {
				boolvar, _ := strconv.ParseBool(*values[param.Name])
				boolMapping[param.Name] = &boolvar
				pflag.BoolVarP(
					boolMapping[param.Name],
					strings.ToLower(param.Name),
					param.ShortName,
					*boolMapping[param.Name],
					param.Description,
				)
			} else {
				pflag.StringVarP(
					values[param.Name],
					strings.ToLower(param.Name),
					param.ShortName,
					*values[param.Name],
					param.Description,
				)
			}
		} else {
			*values[param.Name] = flag.Value.String()
		}
	}
	pflag.Parse()
	for name, value := range boolMapping {
		strValue := strconv.FormatBool(*value)
		values[name] = &strValue
	}
	return values
}
