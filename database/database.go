package database

import (
	"log"
	"os"
	"todos/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	DB *gorm.DB
}

var Database DBinstance

func ConnectToDb() {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to Open Db", err.Error())
		os.Exit(2)
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	db.AutoMigrate(&models.Todo{})
	Database = DBinstance{DB: db}
}
