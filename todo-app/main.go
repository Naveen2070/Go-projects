package main

import (
	"fmt"
	"net/http"
	controllers "todo-app/internals/controller"
	"todo-app/internals/models"
	"todo-app/internals/templates/components/todo"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

const (
	Port = "8080"
)

func main() {
	component := todo.Index(make([]models.Todo, 0))

	router := mux.NewRouter()

	//component endpoints
	router.Handle("/", templ.Handler(component)).Methods("GET")

	// Static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))).Methods("GET")

	// API endpoints
	todoController := controllers.CreateTodoController()

	router.HandleFunc("/api/todos/all", func(w http.ResponseWriter, r *http.Request) {
		todoController.GetTodos(w, r)
	}).Methods("GET")

	router.HandleFunc("/api/todos/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		todoController.GetTodosById(w, r)
	}).Methods("GET")

	router.HandleFunc("/api/todos/add", func(w http.ResponseWriter, r *http.Request) {
		todoController.AddTodo(w, r)
	}).Methods("POST")

	router.HandleFunc("/api/todos/update/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		todoController.UpdateTodo(w, r)
	}).Methods("PUT")

	router.HandleFunc("/api/todos/delete/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		todoController.DeleteTodo(w, r)
	}).Methods("DELETE")

	// Start server
	fmt.Printf("Listening on port %s\n", Port)
	http.ListenAndServe(":"+Port, router)
}
