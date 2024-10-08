package server

import (
	"errors"
	"fmt"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(port string) *Server {
	server := Server{}

	mux := http.NewServeMux()
	mux.HandleFunc("/version", server.VersionHandler)
	mux.HandleFunc("/decode", server.DecodeHandler)
	mux.HandleFunc("/hard-op", server.HardOpHandler)

	server.server = &http.Server{
		Addr:    port,
		Handler: mux,
	}
	return &server
}

func (server *Server) Start() {
	if err := server.server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		fmt.Println(err)
	}
}
