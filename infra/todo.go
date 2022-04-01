package infra

import (
	"gRPC_todo/config"
	"gRPC_todo/entity"
)

type TodoRepository interface {
	Create(*entity.Todo) (*entity.Todo, error)
	GetTodoList() (*entity.Todo, error)
}

type todoRepository struct {
	config.DBHandler
}

func (tr *todoRepository) Create(todo *entity.Todo) (*entity.Todo, error) {
	if err := tr.Conn.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}