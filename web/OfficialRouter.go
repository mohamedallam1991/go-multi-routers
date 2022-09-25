package web

import (
	"fmt"
	"net/http"
)

type OfficialRouterConfig struct {
	Port    string
	Handler http.Handler
}

func newOfficialRouter() *OfficialRouterConfig {
	return &OfficialRouterConfig{}
}

func (o OfficialRouterConfig) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from Official Router, %s!\n", "aa")
}

func (o *OfficialRouterConfig) setRoutes() {
	a := http.NewServeMux()

	http.HandleFunc("/", o.Hello)

	o.Handler = a
}

func (o *OfficialRouterConfig) setPortNum(portNum int) {
	o.Port = fmt.Sprintf(":%v", portNum)
}

func (o *OfficialRouterConfig) getRouterConf() Router {
	return Router{
		Port:    o.Port,
		Handler: o.Handler,
	}
}
