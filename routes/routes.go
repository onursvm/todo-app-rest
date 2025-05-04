package routes

import (
	"net/http"
	"todo-app/controllers"
	"todo-app/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	api.HandleFunc("/todos", controllers.GetToDos).Methods("GET")
	api.HandleFunc("/todos/create", controllers.CreateToDo).Methods("POST")
	api.HandleFunc("/todos/update", controllers.UpdateToDo).Methods("PUT")
	api.HandleFunc("/todos/delete/{id}", controllers.DeleteToDo).Methods("DELETE")

	api.HandleFunc("/steps/create", controllers.CreateStep).Methods("POST")
	api.HandleFunc("/steps/update", controllers.UpdateStep).Methods("PUT")
	api.HandleFunc("/steps/delete/{id}", controllers.DeleteStep).Methods("DELETE")

	return router
}
