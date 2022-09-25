package main

// import (
// 	"net/http"
// )

// // type MMux struct {
// // }

// func (m *MMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		HelloRaw(w, r)
// 	}
// }

// // func HelloRaw(w http.ResponseWriter, req *http.Request) {
// // 	fmt.Fprintf(w, "HelloRaw\n")
// // }

// type Director struct {
// 	builder RouterInterface
// }

// func ANewRouter(b RouterInterface) *Director {
// 	return &Director{
// 		builder: b,
// 	}
// }

// func (d *Director) setBuilder(b RouterInterface) {
// 	d.builder = b
// }

// func (d *Director) buildHouse() http.Handler {
// 	d.builder.setPort()
// 	return d.builder.getRoutes()
// }

// // func Routes() http.Handler {
// // 	chiBuilder := getBuilder("chi")
// // 	// julienBuilder := getBuilder("julien")
// // 	// router := chi.NewRouter()

// // 	// router.Get("/", HelloChi)
// // 	// router.Get("/", HelloChi)

// // 	return chiBuilder.getRoutes()
// // }

// type RouterInterface interface {
// 	setPort()
// 	getRoutes() http.Handler
// }

// func getBuilder(builderType string) RouterInterface {
// 	if builderType == "chi" {
// 		return RouterFromChi()
// 	}

// 	if builderType == "julien" {
// 		return RouterFromJulien()
// 	}
// 	return nil
// }

// //	type House struct {
// //		windowType string
// //		doorType   string
// //		floor      int
// //	}
// // type Router struct {
// // 	config     string
// // 	portNumber string
// // }
