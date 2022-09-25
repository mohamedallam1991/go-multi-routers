package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type ChiConofig struct {
	windowType string
	Port       string
	Handler    http.Handler
}

func newChiRouter() *ChiConofig {
	return &ChiConofig{}
}

func (j ChiConofig) Hello(wa http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(wa, "hello froom cchi, %s!\n", "a")
}

//	func Hello(wa http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(wa, "hello froom cchi, %s!\n", "a")
//	}
func (c *ChiConofig) getRoutes() {
	router := chi.NewRouter()
	router.Get("/", c.Hello)
	router.Get("/hi", c.Hello)

	c.Handler = router
}

// func (b *ChiConofig) ServeHTTP(http.ResponseWriter, *http.Request) {
// 	b.windowType = "Wooden Window"
// }

// func (b *ChiConofig) setDoorType() {
// 	b.doorType = "Wooden Door"
// }

func (c *ChiConofig) SetNumFloor(portNum string) {
	c.Port = portNum
}

func (c *ChiConofig) GetHouse() Router {
	return Router{
		windowType: c.windowType,
		Port:       c.Port,
		Handler:    c.Handler,
	}
}
