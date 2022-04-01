package handler

import "gRPC_todo/infra"

type TodoHandler struct {
	todoRepository *infra.TodoRepository
}

func NewTodoHandler (todoRepository *infra.TodoRepository) *TodoHandler {
	return &TodoHandler{todoRepository: todoRepository}
}

