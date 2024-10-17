package main

import (
	"fmt"
	"time"

	"github.com/kotoproger/home_work_basic/hw10_motion_sensor/sensor"
)

func main() {
	duration := time.Second * 10
	rawData := make(chan int)
	output := make(chan float32)
	defer close(rawData)
	defer close(output)

	readExit := time.After(duration)
	calcExit := time.After(duration + time.Second)
	timer := time.After(duration + time.Second*2)

	go sensor.ReadSensor(rawData, readExit, time.Millisecond*100)
	go sensor.DataProcessor(rawData, output, calcExit)

	for {
		select {
		case data := <-output:
			fmt.Println(data)
		case <-timer:
			return
		}
	}
}
