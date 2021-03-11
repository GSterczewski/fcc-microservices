package services

import (
	"net/http"
)

//WhoamiResponse - response type for whoami miscroservice
type WhoamiResponse struct {
	Ipaddress string `json:"ipaddress"`
	Language  string `json:"language"`
	Software  string `json:"software"`
}

//Whoami - service
type Whoami struct {
	req *http.Request
}

func (wh Whoami) getSoftware() string {
	return wh.req.UserAgent()
}

func (wh Whoami) getLanguage() string {
	return wh.req.Header.Get("Accept-Language")
}

func (wh Whoami) getIP() string {
	return wh.req.RemoteAddr
}

// Parse - main func for Whoami service
func (wh Whoami) Parse() WhoamiResponse {
	r := WhoamiResponse{wh.getIP(), wh.getLanguage(), wh.getSoftware()}
	return r
}

// NewWhoami - constructor
func NewWhoami(req *http.Request) Whoami {
	return Whoami{req}
}
