package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fcc-microservices/router"
	"github.com/fcc-microservices/services"
)

const port = 8080
const apiPath = "/api"

func handleTimestamp(rw http.ResponseWriter, req *http.Request) {
	ts := services.Timestamp{}
	ts.Run(rw)
}

func main() {
	router := router.NewRouter(apiPath)
	router.Register("timestamp/", handleTimestamp)
	router.Init()
	fmt.Printf("FCC-MICROSERVICES run on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
