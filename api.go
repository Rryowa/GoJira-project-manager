package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//start server and anything related

type APIServer struct {
	//port 3000
	addr string
	//db(repository) dependency injection
	repo Repo
}

// Constructor
func NewAPIServer(addr string, newRepo Repo) *APIServer {
	return &APIServer{addr: addr, repo: newRepo}
}

// Serve() initializes a router, register all services(User, Projects, Tasks)
// Listen
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	sub := router.PathPrefix("api/v1").Subrouter()

	//TODO: register services
	tasksService := NewTasksService(s.repo)
	tasksService.RegisterRoutes(router)

	log.Println("Starting API server", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, sub))

}
