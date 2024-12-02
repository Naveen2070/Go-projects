package handlers

import (
	"errors"
	"todo-app/internals/models"
)

type Todo struct {
	Todos models.Todos
}

// TodoHandler interface implementation
type TodoHandler interface {
	GetTodos() []models.Todo
	AddTodo(todo string)
	DeleteTodo(id int)
	UpdateTodo(updatedTodo models.Todo)
	FindTodoById(id int) (models.Todo, error)
}

// Method to get all todos
func (t *Todo) GetTodos() []models.Todo {
	return t.Todos
}

// Method to add a new todo
func (t *Todo) AddTodo(todo string) {
	newTodo := models.Todo{
		ID:   len(t.Todos) + 1,
		Task: todo,
	}
	t.Todos = append(t.Todos, newTodo)
}

// Method to delete a todo by ID
func (t *Todo) DeleteTodo(id int) {
	index := -1
	for i, todo := range t.Todos {
		if todo.ID == id {
			index = i
			break
		}
	}
	if index != -1 {
		t.Todos = append(t.Todos[:index], t.Todos[index+1:]...)
	}
}

// Method to update an existing todo by ID
func (t *Todo) UpdateTodo(updatedTodo models.Todo) {
	for i, todo := range t.Todos {
		if todo.ID == updatedTodo.ID {
			t.Todos[i].Task = updatedTodo.Task
			break
		}
	}
}

// Method to get a todo by ID
func (t *Todo) FindTodoById(id int) (models.Todo, error) {
	for _, todo := range t.Todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return models.Todo{}, errors.New("todo not found")
}
