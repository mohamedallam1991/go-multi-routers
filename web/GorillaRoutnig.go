package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type GorillaRouting struct {
	Port    string
	Handler http.Handler
}

func newGorillaRouter() *GorillaRouting {
	return &GorillaRouting{}
}

func (g *GorillaRouting) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from Gorilla Routing Router, %s!\n", "aa")
}

func (g *GorillaRouting) setRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", g.Hello)
	g.Handler = router
}

func (g *GorillaRouting) setPortNum(portNum int) {
	g.Port = fmt.Sprintf(":%v", portNum)
}

func (g *GorillaRouting) getRouterConf() Router {
	return Router{
		Port:    g.Port,
		Handler: g.Handler,
	}
}
