package main

import (
	"log"
	"net/http"

	"github.com/mohamedallam1991/ticksys/web"
)

const port string = ":8011"

//	func newChiConofig() *ChiConofig {
//	    return &ChiConofig{}
//	}
// type Router struct {
// 	portNumber string
// }

// func (paa Handler) ServeHTTP(builderType string) Handler {

// 	return paa
// }

func main() {
	// server := &http.Server{Addr: "10", Handler: handler}

	// go config()

	// http.Serve()
	// server := &Server{Addr: port, Handler: handler}
	// http.Serve()
	// http.HandleFunc("/", hello)
	// log.Fatal(http.ListenAndServe(port, server))
	// server := &Server{Addr: addr, Handler: handler}
	// return server.ListenAndServe()

	// web.getBuilder

	chiBuilder := web.GetBuilder("chi")
	JshmidtBuilder := web.GetBuilder("Jshmidt")

	director := web.NewRouter(chiBuilder)
	chiRouter := director.BuildHouse(port)
	log.Fatal(http.ListenAndServe(chiRouter.Port, chiRouter.Handler))

	// getRoutes
	// chiRouter.Builder.
	// chiRouter.GetRoutes()
	// a = chiRouter.GetHouse

	// log.Fatal(http.ListenAndServe(chiRouter.Port, director.Builder.GetRoutes()))
	// log.Fatal(http.ListenAndServe(chiRouter.Port, chiRouter.Handler))

	_ = chiRouter
	// fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	// fmt.Printf("Normal House Window Type: %s\n", chiRouter.windowType)
	// fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director = web.NewRouter(JshmidtBuilder)
	Jshmidt := director.BuildHouse(port)
	// Jshmidt.Builder.
	_ = Jshmidt
	log.Fatal(http.ListenAndServe(Jshmidt.Port, Jshmidt.Handler))

	// fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	// fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	// fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}
