package web

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
	d.builder.setRoutes()
	d.builder.setPortNum(portNum)
	return d.builder.getRouterConf()
}
