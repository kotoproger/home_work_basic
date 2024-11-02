package main

import (
	"fmt"
	"os"

	"github.com/kotoproger/home_work_basic/hw12_log_util/log"
)

func main() {
	file, _ := os.Open("/home/kbkondakov/otus/home_work_basic/hw12_log_util/temp.log")

	fmt.Println(
		log.StatCalc(
			log.Filter(
				log.Parse(
					log.Read(file),
				),
				func(record log.Record) bool {
					return record.Level == log.LogLevelInfo
				},
			),
		),
	)
}
