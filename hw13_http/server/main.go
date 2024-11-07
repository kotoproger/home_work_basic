package main

import (
	"fmt"
	"net/http"

	"github.com/kotoproger/home_work_basic/configapp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HTTP %s %s%s from:%s", r.Method, r.Host, r.RequestURI, r.RemoteAddr)
}

func main() {
	c := configapp.ConfigApp{}
	c.AddParam(configapp.ConfigParam{Name: "address", Description: "Address to listen", ShortName: "a", Default: "127.0.0.1"})
	c.AddParam(configapp.ConfigParam{Name: "port", Description: "Port", ShortName: "p", Default: "8880"})

	configapp.GetConfig(c)

	http.HandleFunc("/", handler)
	port, _ := c.GetInt("port")
	http.ListenAndServe(fmt.Sprintf("%s:%d", c.GetString("address"), port), nil)
}
