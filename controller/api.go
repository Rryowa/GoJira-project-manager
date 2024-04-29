package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rryowa/go-jwt-auth/repo"
	"github.com/rryowa/go-jwt-auth/service"
)

//start server and anything related

type APIServer struct {
	//localhost:3000
	addr string
	//db(repository) dependency injection
	repo repo.Repo
}

// Constructor
func NewAPIServer(addr string, newRepo repo.Repo) *APIServer {
	return &APIServer{addr: addr, repo: newRepo}
}

// Initializes a router, register all services(User, Projects, Tasks)
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	sub := router.PathPrefix("/api/v1").Subrouter()

	//TODO: register services
	tasksService := service.NewTasksService(s.repo)
	tasksService.RegisterRoutes(sub)

	log.Println("Starting API server", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, sub))

}
