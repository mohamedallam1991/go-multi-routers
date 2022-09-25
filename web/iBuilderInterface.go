package web

// type Handler interface {
// 	ServeHTTP(http.ResponseWriter, *http.Request)
// }

type IBuilder interface {
	// ServeHTTP(http.ResponseWriter, *http.Request)
	getRoutes()
	SetNumFloor(Theport string)
	GetHouse() Router
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "chi" {
		return newChiRouter()
	}

	if builderType == "Jshmidt" {
		return newJShmidtRouter()
	}
	return nil
}
