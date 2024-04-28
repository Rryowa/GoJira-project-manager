package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type TasksService struct {
	repo Repo
}

func NewTasksService(r Repo) *TasksService {
	return &TasksService{repo: r}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")
}

// we need to access request body
func (ts *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, InvalidRequestError)
		return
	}
	//what if types.go -> Name is empty?

	if err := validateTask(task); err != nil {
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	t, err := ts.repo.CreateTask(task)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, CreateTaskError)
		return
	}
	WriteJSON(w, http.StatusCreated, t)
}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}

func validateTask(task *Task) error {
	if task.Name == "" {
		return NameRequiredError
	}
	if task.ProjectID == 0 {
		return ProjectIDRequiredError
	}
	if task.AssignedTo == 0 {
		return UserIDRequiredError
	}
	return nil
}
