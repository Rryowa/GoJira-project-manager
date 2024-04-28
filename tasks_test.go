package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateTask(t *testing.T) {
	//defining mock

	mr := NewMock()
	service := NewTasksService(mr)

	t.Run("Should return error if name is empty", func(t *testing.T) {
		payload := &Task{
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
		router.HandleFunc("/tasks", service.handleCreateTask)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Error("invalid status code")
		}
	})
}

func testGetTask(t *testing.T) {
	mr := NewMock()
	service := NewTasksService(mr)

	t.Run("Should return task", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/tasks/1", nil)
		// '/' because we creating out small router for testing
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/tasks/{id}", service.handleGetTask)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Error("invalid status code")
		}
	})
}