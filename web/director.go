package web

type Director struct {
	builder IBuilder
}

func NewRouter(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildRouter(portNum int) Router {
	d.builder.setRoutes()
	d.builder.setPortNum(portNum)
	return d.builder.getRouterConf()
}
