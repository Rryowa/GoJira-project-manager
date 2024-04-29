package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rryowa/go-jwt-auth/entity"
	"github.com/rryowa/go-jwt-auth/repo"
	"github.com/rryowa/go-jwt-auth/util"
)

type TasksService struct {
	repo repo.Repo
}

func NewTasksService(r repo.Repo) *TasksService {
	return &TasksService{repo: r}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.HandleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.HandleGetTask).Methods("GET")
}

// we need to access request body
func (ts *TasksService) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	var task *entity.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, util.InvalidRequestError)
		return
	}

	//what if types.go -> Name is empty?
	if err := validateTask(task); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	t, err := ts.repo.CreateTask(task)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, util.CreateTaskError)
		return
	}
	util.WriteJSON(w, http.StatusCreated, t)
}

func (ts *TasksService) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		util.WriteJSON(w, http.StatusBadRequest, util.IdRequiredError)
		return
	}

	t, err := ts.repo.GetTask(id)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, util.TaskRequired)
		return
	}

	util.WriteJSON(w, http.StatusOK, t)
}

func validateTask(task *entity.Task) error {
	if task.Name == "" {
		return util.NameRequiredError
	}
	if task.ProjectID == 0 {
		return util.ProjectIDRequiredError
	}
	if task.AssignedTo == 0 {
		return util.UserIDRequiredError
	}
	return nil
}
