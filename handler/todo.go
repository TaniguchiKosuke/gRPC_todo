package handler

import (
	"context"
	"gRPC_todo/entity"
	"gRPC_todo/infra"
	"gRPC_todo/proto"
	"log"

	"github.com/google/uuid"
)

type TodoHandler struct {
	proto.UnimplementedTodoCommandServer
	proto.UnimplementedTodoQueryServer

	todoRepository infra.TodoRepository
}

func NewTodoHandler(todoRepository infra.TodoRepository) *TodoHandler {
	return &TodoHandler{todoRepository: todoRepository}
}

func (th *TodoHandler) Create(ctx context.Context, req *proto.TodoCreateRequest) (*proto.TodoCreateResponse, error) {
	todoUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	todo := &entity.Todo{
		ID:    todoUUID.String(),
		Title: req.Title,
		Body:  req.Body,
	}

	createdTodo, err := th.todoRepository.Create(todo)
	if err != nil {
		return nil, err
	}
	log.Println("create todo called !!!!!")

	return &proto.TodoCreateResponse{
		Id: createdTodo.ID,
	}, nil
}

func (th *TodoHandler) Get(ctx context.Context, req *proto.TodoGetRequest) (*proto.TodoGetResponse, error) {
	todo, err := th.todoRepository.GetTodoByID(req.Id)
	if err != nil {
		return nil, err
	}

	resTodo := &proto.TodoGetResponse_Todo{
		Id:    todo.ID,
		Title: todo.Title,
		Body:  todo.Body,
	}
	return &proto.TodoGetResponse{
		Item: resTodo,
	}, nil
}
