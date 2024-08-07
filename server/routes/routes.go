package routes

import (
	"github.com/unamdev0/go-todo-app/middlewares"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", middlewares.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middlewares.CreateTask).Methods("POST","OPTIONS")
	
	router.HandleFunc("/api/tasks/{id}", middlewares.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task/undo/{id}", middlewares.UndoDelete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task/delete/{id}", middlewares.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/task/delete/all", middlewares.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
