package handlers

import (
	"net/http"
	"todo-app/internal/models"
	"todo-app/internal/templates"

	"github.com/a-h/templ"
)

var nextID = 1

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		task := r.FormValue("task")
		models.Todos = append(models.Todos, models.Todo{ID: nextID, Task: task})
		nextID++
	}
	templ.Render(w, templates.TodoItem(models.Todos[len(models.Todos)-1]))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templ.Render(w, templates.Index(models.Todos))
}
