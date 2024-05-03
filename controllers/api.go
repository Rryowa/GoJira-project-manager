package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rryowa/Gojira-project-manager/repo"
	"github.com/rryowa/Gojira-project-manager/services"
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

	usersService := services.NewUserService(s.repo)
	usersService.RegisterRoutes(sub)
	tasksService := services.NewTasksService(s.repo)
	tasksService.RegisterRoutes(sub)

	log.Println("Starting API server", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, sub))

}
