package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rryowa/Gojira-project-manager/entity"
	"github.com/rryowa/Gojira-project-manager/repo"
	"github.com/rryowa/Gojira-project-manager/utils"
)

type TasksService struct {
	repo repo.Repo
}

func NewTasksService(r repo.Repo) *TasksService {
	return &TasksService{repo: r}
}

func (ts *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", utils.WithJWTAuth(ts.HandleCreateTask, ts.repo)).Methods("POST")
	r.HandleFunc("/tasks/{id}", utils.WithJWTAuth(ts.HandleGetTask, ts.repo)).Methods("GET")
}

// we need to access request body
func (ts *TasksService) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var task *entity.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest,
			utils.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	//what if types.go -> Name is empty?
	if err := validateTask(task); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest,
			utils.ErrorResponse{Error: err.Error()})
		return
	}

	t, err := ts.repo.CreateTask(task)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError,
			utils.ErrorResponse{Error: "Error creating task"})
		return
	}
	utils.WriteJSON(w, http.StatusCreated, t)
}

func (ts *TasksService) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		utils.WriteJSON(w, http.StatusBadRequest,
			utils.ErrorResponse{Error: "Error id required"})
		return
	}

	t, err := ts.repo.GetTask(id)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError,
			utils.ErrorResponse{Error: "Error there is no task"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, t)
}

func validateTask(task *entity.Task) error {
	if task.Name == "" {
		return utils.ErrNameRequired
	}
	if task.ProjectID == 0 {
		return utils.ErrProjectIDRequired
	}
	if task.AssignedTo == 0 {
		return utils.ErrUserIDRequired
	}
	return nil
}
