package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type ChiConofig struct {
	Port    string
	Handler http.Handler
}

func newChiRouter() *ChiConofig {
	return &ChiConofig{}
}

func (c *ChiConofig) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from root of %s Router!\n", "Chi")
	// fmt.Fprintf(wa, "hello froom cchi, %s!\n", "a")
}

func (c *ChiConofig) setRoutes() {
	router := chi.NewRouter()
	router.Get("/", c.Hello)

	c.Handler = router
}

func (c *ChiConofig) setPortNum(portNum int) {
	c.Port = fmt.Sprintf(":%v", portNum)
}

func (c *ChiConofig) getRouterConf() Router {
	return Router{
		Port:    c.Port,
		Handler: c.Handler,
	}
}
