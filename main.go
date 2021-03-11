package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fcc-microservices/config"
	"github.com/fcc-microservices/responder"
	"github.com/fcc-microservices/router"
	"github.com/fcc-microservices/services"
	"github.com/fcc-microservices/static"
)

func getDateStringFromURI(req *http.Request, path string) string {
	input := req.URL.Path[len(path):]
	if len(input) == 0 {
		return ""
	}
	inputSlice := strings.Split(input, "/")
	return inputSlice[0]

}
func handleTimestamp(rw http.ResponseWriter, req *http.Request) {
	ts := services.Timestamp{}
	responder := responder.NewResponder(rw)
	timestamp, err := ts.Parse(getDateStringFromURI(req, config.TimestampServicePath))
	if err != nil {
		fmt.Fprintf(rw, "Error : %s", err)
	}
	responder.ServeJSON(&timestamp)

}

func handleHomePage(rw http.ResponseWriter, req *http.Request) {
	responder := responder.NewResponder(rw)
	wd, _ := os.Getwd()

	timestampServiceLink := static.Link{Name: "Timestamp service", Href: config.TimestampServicePath}
	links := []static.Link{timestampServiceLink}
	pd := static.PageData{Title: config.PageTitle, Links: links}
	responder.ServeHTML(fmt.Sprintf("%s/static/layout.html", wd), pd)
}
func main() {
	router := router.NewRouter()
	router.Register(config.TimestampServicePath, handleTimestamp)
	router.Register(config.HomePath, handleHomePage)
	router.Init()
	fmt.Printf("FCC-MICROSERVICES run on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil))
}
