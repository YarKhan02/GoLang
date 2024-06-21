package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	store Store
}

// constructor
func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store} 
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("Starting the api server at: ", s.addr)
	log.Println(http.ListenAndServe(s.addr, subrouter))
}