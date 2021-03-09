package router

import (
	"fmt"
	"net/http"
)

//Router - struct responsible for registering and handling routes
type Router struct {
	mainPath string
	handlers map[string]http.HandlerFunc
}

//Register - handler registration
func (r *Router) Register(path string, handler http.HandlerFunc) {
	p := fmt.Sprintf("%s/%s", r.mainPath, path)
	r.handlers[p] = handler
}

//Init - after routes registration use this method to init Router
func (r Router) Init() {
	for path, handler := range r.handlers {
		http.HandleFunc(path, handler)
	}
}

//NewRouter - simple Router creator function
func NewRouter(mainPath string) Router {
	handlers := make(map[string]http.HandlerFunc)
	return Router{mainPath, handlers}
}
