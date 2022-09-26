package main

import (
	"log"
	"net/http"

	"github.com/mohamedallam1991/ticksys/web"
)

const port int = 8011

func main() {

	setup := web.NewRouter(web.GetBuilder(web.Routers.Gorilla))
	router := setup.BuildRouter(port)
	log.Fatal(http.ListenAndServe(router.Port, router.Handler))
}
