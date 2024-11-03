package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kotoproger/home_work_basic/hw12_log_util/log"
	"github.com/kotoproger/home_work_basic/hw12_log_util/settings"
)

func main() {
	settings := settings.ParseParams(
		settings.ParseEnv(
			settings.Settings{LogLevel: log.LogLevelInfo},
		),
	)

	if settings.InputFileName == "" {
		panic("Путь до лог файла не задан")
	}

	file, err := os.Open(settings.InputFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	statistics := log.StatCalc(
		log.Filter(
			log.Parse(
				log.Read(file),
			),
			func(record log.Record) bool {
				return record.Level == settings.LogLevel
			},
		),
	)

	jsonData, _ := json.Marshal(statistics)
	if settings.OutputFileName == "" {
		fmt.Println(string(jsonData))
	} else {
		output, outerr := os.OpenFile(settings.OutputFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
		if outerr != nil {
			panic(outerr)
		}
		defer output.Close()
		output.Write(jsonData)
	}
}
