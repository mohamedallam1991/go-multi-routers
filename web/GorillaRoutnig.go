package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type GorillaRouting struct {
	windowType string
	Port       string
	Handler    http.Handler
}

func newGorillaRouter() *GorillaRouting {
	return &GorillaRouting{}
}

func (o GorillaRouting) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from Gorilla Routing Router, %s!\n", "aa")
}

func (o *GorillaRouting) setRoutes() {
	// router := httprouter.New()
	// router := http.Handl("/", o.Hello)
	router := mux.NewRouter()
	// http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "hello from J Shmidt, %s!\n", "aa")
	// })
	router.HandleFunc("/", o.Hello)

	// router.GET("/", j.Hello)

	o.Handler = router
}

func (o *GorillaRouting) setPortNum(portNum string) {
	o.Port = portNum
}

func (o *GorillaRouting) getRouterConf() Router {
	return Router{
		windowType: o.windowType,
		Port:       o.Port,
		Handler:    o.Handler,
	}
}
