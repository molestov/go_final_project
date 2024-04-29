package database

import (
	"errors"
	"log"
	"os"

	"github.com/molestov/go_final_project/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DB_QUERY_LIMIT = 30

var Db *gorm.DB

func getDBPath() string {
	const DBFile = "scheduler.db"
	if val, exists := os.LookupEnv("TODO_DBFILE"); exists {
		return val
	}
	return DBFile
}

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open(getDBPath()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database.\n", err)
		os.Exit(1)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migration")
	db.AutoMigrate(&models.Task{})

	Db = db
}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := Db.Order("date").Limit(DB_QUERY_LIMIT).Find(&tasks)
	return tasks, result.Error
}

func GetTaskById(id uint) (models.Task, error) {
	var tasks models.Task
	result := Db.Limit(1).Find(&tasks, models.Task{ID: id})
	if result.RowsAffected == 0 {
		return tasks, errors.New("no record found")
	}
	return tasks, result.Error
}
