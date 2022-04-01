package server

import (
	"gRPC_todo/config"
	"gRPC_todo/handler"
	"gRPC_todo/infra"
)

func InjectTodoAPI() *handler.TodoHandler {
	dbHandler := config.NewDBHandler()
	todoRepository := infra.NewTodoRepository(*dbHandler)
	todoHandler := handler.NewTodoHandler(todoRepository)

	return todoHandler
}
