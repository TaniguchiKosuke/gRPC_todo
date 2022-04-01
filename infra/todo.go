package infra

import (
	"gRPC_todo/config"
	"gRPC_todo/entity"
)

type TodoRepository interface {
	Create(*entity.Todo) (*entity.Todo, error)
	GetTodoByID(string) (*entity.Todo, error)
}

type todoRepository struct {
	config.DBHandler
}

func NewTodoRepository(dbHandler config.DBHandler) TodoRepository {
	return &todoRepository{DBHandler: dbHandler}
}

func (tr *todoRepository) Create(todo *entity.Todo) (*entity.Todo, error) {
	if err := tr.Conn.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func (tr *todoRepository) GetTodoByID(ID string) (*entity.Todo, error) {
	var todo *entity.Todo
	if err := tr.Conn.Where("id = ?", ID).Find(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}
