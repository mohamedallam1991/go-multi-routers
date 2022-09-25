package web

import "net/http"

type Router struct {
	Port    string
	Handler http.Handler
}
