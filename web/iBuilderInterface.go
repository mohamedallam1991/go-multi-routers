package web

import "log"

type IBuilder interface {
	setRoutes()
	setPortNum(Theport string)
	getRouterConf() Router
}

func GetBuilder(builderType string) IBuilder {

	switch builderType {
	case "chi":
		return newChiRouter()
	case "Jshmidt":
		return newJShmidtRouter()
	case "official":
		return newOfficialRouter()
	case "gorilla":
		return newGorillaRouter()
	default:
		log.Fatal("router name is incorrect")
	}
	// if builderType == "chi" {
	// 	return newChiRouter()
	// }

	// if builderType == "Jshmidt" {
	// 	return newJShmidtRouter()
	// }

	// if builderType == "Jshmidt" {
	// 	return soomethinig()
	// }
	return nil
}
