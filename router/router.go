package router

import "net/http"

type Router struct {
	routes map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) Register(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) GetHandler(path string) http.HandlerFunc {
	return r.routes[path]
}
