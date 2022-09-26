package web

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type JShmidtConfig struct {
	windowType string
	Port       string
	Handler    http.Handler
}

func newJShmidtRouter() *JShmidtConfig {
	return &JShmidtConfig{}
}

func (j JShmidtConfig) Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello from root of %s Router!\n", "JulienSchmidt")

	// fmt.Fprintf(w, "hello from J Shmidt, %s!\n", ps.ByName("name"))
}

func (j *JShmidtConfig) setRoutes() {
	router := httprouter.New()
	router.GET("/", j.Hello)

	j.Handler = router
}

func (j *JShmidtConfig) setPortNum(portNum int) {
	j.Port = fmt.Sprintf(":%v", portNum)
}

func (j *JShmidtConfig) getRouterConf() Router {
	return Router{
		Port:    j.Port,
		Handler: j.Handler,
	}
}
