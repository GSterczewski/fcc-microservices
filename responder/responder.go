package responder

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/fcc-microservices/static"
)

// Responder - general struct for serving responses to the client
type Responder struct {
	rw http.ResponseWriter
}

// InternalServerError - writes status 500 and optional message to the client, when error occurs
func (s Responder) InternalServerError(message string) {
	s.rw.WriteHeader(http.StatusInternalServerError)
	s.rw.Write([]byte(message))

}

// ServeJSON - write json response to the client
func (s Responder) ServeJSON(payload interface{}) {

	json, err := json.Marshal(payload)
	if err != nil {
		s.InternalServerError("")
		return
	}
	s.rw.Header().Set("Content-Type", "application/json")
	s.rw.WriteHeader(http.StatusOK)
	s.rw.Write(json)
}

// ServeHTML - serves specified html template
func (s Responder) ServeHTML(path string, pd static.PageData) {
	s.rw.Header().Set("Content-Type", "text/html")
	s.rw.WriteHeader(http.StatusOK)

	html, parsingErr := template.ParseFiles(path)
	if parsingErr != nil {
		fmt.Println(parsingErr)
		return
	}

	err := html.Execute(s.rw, pd)
	if err != nil {
		log.Println(err)
		return
	}
}

// NewResponder - constructor function
func NewResponder(rw http.ResponseWriter) Responder {
	return Responder{rw}
}
