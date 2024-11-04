package settings

import (
	"github.com/kotoproger/home_work_basic/hw12_log_util/log"
	"github.com/spf13/pflag"
)

func ParseParams(settings Settings) Settings {
	pflag.StringVarP(&settings.InputFileName, "file", "f", settings.InputFileName, "Путь к лог файлу")
	var level string
	pflag.StringVarP(&level, "level", "l", string(settings.LogLevel), "Уровень логов для анализа")
	pflag.StringVarP(
		&settings.OutputFileName,
		"output",
		"o",
		settings.OutputFileName,
		"Путь к файлу, в который будет записана статистика (необязательный флаг, если не"+
			" указан, статистика выводится в стандартный поток вывода)",
	)
	pflag.Parse()

	if level != "" {
		settings.LogLevel = log.Level(level)
	}

	return settings
}
