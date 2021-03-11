package router

import (
	"net/http"
)

//Router - struct responsible for registering and handling routes
type Router struct {
	handlers map[string]http.HandlerFunc
}

//Register - handler registration
func (r *Router) Register(path string, handler http.HandlerFunc) {

	r.handlers[path] = handler
}

//Init - after routes registration use this method to init Router
func (r Router) Init() {
	for path, handler := range r.handlers {
		http.HandleFunc(path, handler)
	}
}

//NewRouter - simple Router creator function
func NewRouter() Router {
	handlers := make(map[string]http.HandlerFunc)
	return Router{handlers}
}
