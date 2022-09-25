package web

import "net/http"

type Router struct {
	windowType string
	Port       string
	Handler    http.Handler
}
