package web

import (
	"log"
)

type IBuilder interface {
	setRoutes()
	setPortNum(portNum int)
	getRouterConf() Router
}

var Routers = newRouterRegistry()

func newRouterRegistry() *routerRegistry {
	return &routerRegistry{
		Chi:      "chi",
		Gorilla:  "gorilla",
		JShmidt:  "jshmidt",
		Official: "offiicial",
	}
}

type routerRegistry struct {
	Chi      string
	Gorilla  string
	JShmidt  string
	Official string
}

func GetBuilder(builderType string) IBuilder {

	switch builderType {
	case newRouterRegistry().Chi:
		return newChiRouter()
	case newRouterRegistry().JShmidt:
		return newJShmidtRouter()
	case newRouterRegistry().Official:
		return newOfficialRouter()
	case newRouterRegistry().Gorilla:
		return newGorillaRouter()
	default:
		log.Fatal("router name is incorrect")
	}
	return nil
}
