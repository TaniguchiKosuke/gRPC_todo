package config

import (
	"gRPC_todo/entity"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBHandler struct {
	Conn *gorm.DB
}

func NewDBHandler() *DBHandler {
	conn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}
	conn.AutoMigrate(&entity.Todo{})

	sqlHandler := new(DBHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}
