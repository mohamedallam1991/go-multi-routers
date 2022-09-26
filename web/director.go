package web

import "log"

type Director struct {
	builder IBuilder
}

func NewRouter(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// Build set builder, to the builder we pass ini
func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

// Build the routes, takes the port number as a parameter
func (d *Director) BuildRouter(portNum int) Router {
	if portNum < 1500 {
		log.Fatal("the port number should be from 2000 all the way up")
	}
	d.builder.setRoutes()
	d.builder.setPortNum(portNum)
	return d.builder.getRouterConf()
}
