package web

type Director struct {
	Builder IBuilder
}

func NewRouter(B IBuilder) *Director {
	return &Director{
		Builder: B,
	}
}

func (d *Director) setBuilder(B IBuilder) {
	d.Builder = B
}

func (d *Director) BuildRouter(Theport string) Router {
	d.Builder.setRoutes()
	d.Builder.setPortNum(Theport)
	return d.Builder.getRouterConf()
}
