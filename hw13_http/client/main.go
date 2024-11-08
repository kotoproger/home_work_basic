package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/kotoproger/home_work_basic/configapp"
)

func main() {
	c := configapp.ConfigApp{}
	c.AddParam(configapp.ConfigParam{
		Name: "host", Description: "Server host with port", ShortName: "h", Default: "127.0.0.1:8880",
	})
	c.AddParam(configapp.ConfigParam{
		Name: "uri", Description: "Uri", ShortName: "u", Default: "/",
	})
	c.AddParam(configapp.ConfigParam{
		Name: "method", Description: "Method", ShortName: "m", Default: "GET",
	})

	configapp.GetConfig(c)
	client := http.Client{}
	uri, err := url.Parse(fmt.Sprintf("http://%s%s", c.GetString("host"), c.GetString("uri")))
	if err != nil {
		panic(err)
	}
	request := http.Request{Method: c.GetString("method"), URL: uri}
	response, requestErr := client.Do(&request)
	if requestErr != nil {
		panic(requestErr)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка чтения", err)
		return
	}

	fmt.Println(string(body))
}
