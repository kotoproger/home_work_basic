package main

import (
	"fmt"
	"time"

	"github.com/kotoproger/home_work_basic/hw10_motion_sensor/sensor"
)

func main() {
	duration := time.Second * 10
	readExit := time.After(duration)

	output := sensor.DataProcessor(
		sensor.ReadSensor(readExit, time.Millisecond*100),
	)

	for data := range output {
		fmt.Println(data)
	}
}
