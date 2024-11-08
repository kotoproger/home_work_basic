package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kotoproger/home_work_basic/configapp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	answer := fmt.Sprintf(
		"HTTP %s %s%s from:%s with body:%s", r.Method, r.Host, r.RequestURI, r.RemoteAddr, r.Body,
	)
	fmt.Fprint(w, answer)
	fmt.Println(answer)
}

func main() {
	c := configapp.ConfigApp{}
	c.AddParam(configapp.ConfigParam{
		Name: "address", Description: "Address to listen", ShortName: "a", Default: "127.0.0.1",
	})
	c.AddParam(configapp.ConfigParam{
		Name: "port", Description: "Port", ShortName: "p", Default: "8880",
	})

	configapp.GetConfig(c)

	http.HandleFunc("/", handler)
	port, _ := c.GetInt("port")

	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", c.GetString("address"), port),
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
