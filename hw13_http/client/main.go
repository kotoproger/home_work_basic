package main

import (
	"fmt"
	"net"

	"github.com/kotoproger/home_work_basic/configapp"
)

func main() {
	c := configapp.ConfigApp{}
	c.AddParam(configapp.ConfigParam{Name: "host", Description: "Server host with port", ShortName: "h", Default: "127.0.0.1:8880"})
	c.AddParam(configapp.ConfigParam{Name: "uri", Description: "Uri", ShortName: "u", Default: "/"})
	c.AddParam(configapp.ConfigParam{Name: "method", Description: "Method", ShortName: "m", Default: "GET"})

	configapp.GetConfig(c)

	conn, err := net.Dial("tcp", c.GetString("host"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := fmt.Sprintf(
		"%s %s HTTP/1.1\r\nHost: %s\r\n\r\n",
		c.GetString("method"),
		c.GetString("uri"),
		c.GetString("host"),
	)
	conn.Write([]byte(message))

	// Получаем ответ от сервера
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:n]))
}
