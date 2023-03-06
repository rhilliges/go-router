package router

import (
	"net/http"
	"strings"
)

type Router struct {
	routes map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) Register(path string, handler http.HandlerFunc) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	r.routes[path] = handler
}

func (r *Router) GetHandler(path string) http.HandlerFunc {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	parts := strings.Split(path, "/")
	for k, handler := range r.routes {
		if strings.Split(k, "/")[0] == parts[0] {
			return handler
		}
	}
	return nil
}

// ServeHTTP implements http.Handler
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler := r.routes[request.URL.Path]
	if handler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, request)
}
