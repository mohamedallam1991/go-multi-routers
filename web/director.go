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

func (d *Director) BuildHouse(Theport string) Router {
	// d.builder.setDoorType()
	// d.builder.setWindowType()
	d.Builder.getRoutes()
	d.Builder.SetNumFloor(Theport)
	return d.Builder.GetHouse()
}
