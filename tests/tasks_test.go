package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rryowa/Gojira-project-manager/entity"
	"github.com/rryowa/Gojira-project-manager/services"
)

// POST/tasks
func TestCreateTask(t *testing.T) {
	//defining mock
	memoryRepo := &MockRepo{}
	// now memoryRepo using Repo interface
	// so we can create new task
	service := services.NewTasksService(memoryRepo)
	t.Run("Should return error if name is empty", func(t *testing.T) {
		payload := &entity.Task{
			Name: "",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		// '/' because we creating out small router for testing
		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/tasks", service.HandleCreateTask)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Error("invalid status code")
		}
	})
}

func TestGetTask(t *testing.T) {
	memoryRepo := &MockRepo{}
	service := services.NewTasksService(memoryRepo)

	t.Run("Should return task", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/tasks/1", nil)
		// '/' because we creating out small router for testing
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/tasks/{id}", service.HandleGetTask)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Error("invalid status code")
		}
	})
}
