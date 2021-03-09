package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fcc-microservices/router"
	"github.com/fcc-microservices/services"
)

const port = 8080
const apiPath = "/api"

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
	timestamp, err := ts.Parse(getDateStringFromURI(req, "/api/timestamp/"))
	if err != nil {
		fmt.Fprintf(rw, "Error : %s", err)
	}
	encoder := json.NewEncoder(rw)
	encoder.Encode(&timestamp)
	fmt.Fprint(rw)

}

func main() {
	router := router.NewRouter(apiPath)
	router.Register("timestamp/", handleTimestamp)
	router.Init()
	fmt.Printf("FCC-MICROSERVICES run on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
