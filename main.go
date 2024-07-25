package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	InProgress TaskStatus = "in-progress"
	Completed  TaskStatus = "completed"
)

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
}

var tasks = make(map[string]Task)
var mu sync.Mutex

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthCheck).Methods("GET")
	r.Use(RequestLoggerMiddleware(r))
	log.Fatal(http.ListenAndServe(":8080", r))
}

// RequestLoggerMiddleware is a middleware function that logs the details of each incoming HTTP request.
func RequestLoggerMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				log.Printf(
					"[%s] %s %s %s",
					req.Method,
					req.Host,
					req.URL.Path,
					req.URL.RawQuery,
				)
			}()
			next.ServeHTTP(w, req)
		})
	}
}

// healthCheck is an HTTP handler function that checks the health of the server.
//
// It returns I'm healthy message
func healthCheck(w http.ResponseWriter, r *http.Request) {
	message := "I'm healthy"
	response := map[string]string{
		"message": message,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func createTask(w http.ResponseWriter, r *http.Request) {

}

func getTasks(w http.ResponseWriter, r *http.Request) {

}

func getTask(w http.ResponseWriter, r *http.Request) {

}

func updateTask(w http.ResponseWriter, r *http.Request) {

}

func deleteTask(w http.ResponseWriter, r *http.Request) {

}
