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

type muxer struct{}

func (s *muxer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	switch url {
	case "/":
		fmt.Fprintf(w, "hello from root of %s Router!\n", "Official")
	default:
		http.NotFound(w, r)
	}
}

func (o *OfficialRouterConfig) setRoutes() {
	a := &muxer{}

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
