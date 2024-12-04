package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo-app/internals/models"
	handlers "todo-app/internals/service"
	todoItems "todo-app/internals/templates/components/todo/partials"
)

// Global Todos instance
var Todos = &handlers.Todo{
	Todos: []models.Todo{
		{ID: 1, Task: "Task 1"},
		{ID: 2, Task: "Task 2"},
	},
}

var TodoHandler = handlers.TodoHandler(Todos)

type TodoController struct{}

func CreateTodoController() *TodoController {
	return &TodoController{}
}

// getTodos handles GET requests to retrieve todos
func (c *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := TodoHandler.GetTodos()

	component := todoItems.TodoItems(data)

	component.Render(r.Context(), w)
}

// getTodosById handles GET requests to retrieve a todo by ID
func (c *TodoController) GetTodosById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stringId := strings.Split(r.URL.Path, "/")[3]
	id, _ := strconv.ParseInt(stringId, 10, 64)

	if id >= 0 {
		todo, err := TodoHandler.FindTodoById(int(id))
		if err != nil {
			http.Error(w, "Todo not found", http.StatusNotFound)
			return
		}

		data, err := json.Marshal(todo)
		if err != nil {
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}

		w.Write(data)
		return
	}
}

// addTodo handles POST requests to add a new todo
func (c *TodoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payload struct {
		Task string `json:"task"`
	}

	payload.Task = r.FormValue("task")

	if payload.Task == "" {
		http.Error(w, "Task is required", http.StatusBadRequest)
		return
	}

	TodoHandler.AddTodo(payload.Task)
	w.WriteHeader(http.StatusCreated)
	c.GetTodos(w, r)
}

// deleteTodo handles DELETE requests to remove a todo by ID
func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.ParseInt(strings.Split(r.URL.Path, "/")[4], 10, 64)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	todo, err := TodoHandler.FindTodoById(int(id))
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	TodoHandler.DeleteTodo(int(todo.ID))
	w.WriteHeader(http.StatusOK)
	// c.GetTodos(w, r)
}

// updateTodo handles PUT requests to update an existing todo
func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.ParseInt(strings.Split(r.URL.Path, "/")[4], 10, 64)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	var payload struct {
		Task string `json:"task"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if payload.Task == "" {
		http.Error(w, "Task is required", http.StatusBadRequest)
		return
	}

	todo, err := TodoHandler.FindTodoById(int(id))
	if err != nil {
		fmt.Println(err, "Todo not found")
		fmt.Println(todo)
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	updateTodo := models.Todo{
		ID:   int(id),
		Task: payload.Task,
	}

	TodoHandler.UpdateTodo(updateTodo)
	w.WriteHeader(http.StatusOK)

	updatedTodo := []models.Todo{
		{ID: int(id), Task: payload.Task},
	}

	component := todoItems.TodoItems(updatedTodo)

	component.Render(r.Context(), w)
}
