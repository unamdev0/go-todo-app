package routes

import (
	"go-todo-app/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task/undo/{id}", middleware.UndoDelete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/task/delete/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/task/delete/all", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
