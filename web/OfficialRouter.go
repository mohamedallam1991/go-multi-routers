package web

import (
	"fmt"
	"net/http"
)

type OfficialRouterConfig struct {
	windowType string
	Port       string
	Handler    http.Handler
}

func newOfficialRouter() *OfficialRouterConfig {
	return &OfficialRouterConfig{}
}

func (o OfficialRouterConfig) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from Official Router, %s!\n", "aa")
}

func (o *OfficialRouterConfig) setRoutes() {
	// router := httprouter.New()
	// router := http.Handl("/", o.Hello)
	a := http.NewServeMux()
	// http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "hello from J Shmidt, %s!\n", "aa")
	// })
	http.HandleFunc("/", o.Hello)

	// router.GET("/", j.Hello)

	o.Handler = a
}

func (o *OfficialRouterConfig) setPortNum(portNum string) {
	o.Port = portNum
}

func (o *OfficialRouterConfig) getRouterConf() Router {
	return Router{
		windowType: o.windowType,
		Port:       o.Port,
		Handler:    o.Handler,
	}
}
