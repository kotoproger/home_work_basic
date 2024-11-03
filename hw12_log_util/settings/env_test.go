package settings

import (
	"os"
	"testing"

	"github.com/kotoproger/home_work_basic/hw12_log_util/log"
	"github.com/stretchr/testify/assert"
)

func TestParseEnv(t *testing.T) {
	testCases := []struct {
		name   string
		sets   map[string]string
		input  Settings
		output Settings
	}{
		{
			name:   "without env",
			sets:   make(map[string]string),
			input:  Settings{InputFileName: "some filename", LogLevel: log.LogLevelFatal, OutputFileName: "some output"},
			output: Settings{InputFileName: "some filename", LogLevel: log.LogLevelFatal, OutputFileName: "some output"},
		},
		{
			name:   "with file",
			sets:   map[string]string{"LOG_ANALYZER_FILE": "some new file name"},
			input:  Settings{InputFileName: "some filename", LogLevel: log.LogLevelFatal, OutputFileName: "some output"},
			output: Settings{InputFileName: "some new file name", LogLevel: log.LogLevelFatal, OutputFileName: "some output"},
		},
		{
			name:   "with level",
			sets:   map[string]string{"LOG_ANALYZER_LEVEL": "debug"},
			input:  Settings{InputFileName: "some filename", LogLevel: log.LogLevelFatal, OutputFileName: "some output"},
			output: Settings{InputFileName: "some filename", LogLevel: log.LogLevelDebug, OutputFileName: "some output"},
		},
		{
			name:   "with output",
			sets:   map[string]string{"LOG_ANALYZER_OUTPUT": "some new output"},
			input:  Settings{InputFileName: "some filename", LogLevel: log.LogLevelFatal, OutputFileName: "some output"},
			output: Settings{InputFileName: "some filename", LogLevel: log.LogLevelFatal, OutputFileName: "some new output"},
		},
		{
			name: "with all",
			sets: map[string]string{
				"LOG_ANALYZER_OUTPUT": "some new output",
				"LOG_ANALYZER_LEVEL":  "debug",
				"LOG_ANALYZER_FILE":   "some new file name",
			},
			input: Settings{InputFileName: "some filename", LogLevel: log.LogLevelFatal, OutputFileName: "some output"},
			output: Settings{
				InputFileName:  "some new file name",
				LogLevel:       log.LogLevelDebug,
				OutputFileName: "some new output",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				os.Clearenv()
				for name, value := range tc.sets {
					os.Setenv(name, value)
				}
				assert.Equal(t, tc.output, ParseEnv(tc.input))
			},
		)
	}
}
