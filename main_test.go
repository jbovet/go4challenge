package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHealthCheck(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthCheck).Methods("GET")
	req := httptest.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedMessage := "{\"message\":\"I'm healthy\"}"
	if rr.Body.String() != expectedMessage {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), expectedMessage)
	}
}

func TestCreateTask(t *testing.T) {
	t.Error("should be implemented")
}

func TestGetTasks(t *testing.T) {
	t.Error("should be implemented")
}

func TestGetTask(t *testing.T) {
	t.Error("should be implemented")
}

func TestUpdateTask(t *testing.T) {
	t.Error("should be implemented")
}

func TestDeleteTask(t *testing.T) {
	t.Error("should be implemented")
}
