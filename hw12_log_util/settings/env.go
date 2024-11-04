package settings

import (
	"os"

	"github.com/kotoproger/home_work_basic/hw12_log_util/log"
)

func ParseEnv(settings Settings) Settings {
	fileName, ok := os.LookupEnv("LOG_ANALYZER_FILE")
	if ok {
		settings.InputFileName = fileName
	}

	level, ok := os.LookupEnv("LOG_ANALYZER_LEVEL")
	if ok {
		settings.LogLevel = log.Level(level)
	}

	output, ok := os.LookupEnv("LOG_ANALYZER_OUTPUT")
	if ok {
		settings.OutputFileName = output
	}

	return settings
}
