package server

import (
	"net/http"
)

func (server *Server) VersionHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		type VersionResponse struct {
			Version string `json:"version"`
		}
		server.Answer(response, &VersionResponse{"0.1.0"})
	} else {
		http.Error(response, "Only GET Method of version is allowed",
			http.StatusMethodNotAllowed)
	}
}
