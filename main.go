package main

import (
	"fmt"
	"os"

	"github.com/gitnoober/grawler/queue"
	"github.com/gitnoober/grawler/repository"
	"github.com/gitnoober/grawler/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
const (
	BUFFER_SIZE = 1000
	NUM_WORKERS = 10
)

func InitDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}


func main() {
	db, err := InitDB()
	if err != nil {
		panic("failed to connect database")
	}
	taskRepo := repository.NewTaskRepository(db)
	urlRepo := repository.NewUrlRepository(db)
	taskResponseRepo := repository.NewTaskResponseRepository(db)
	queue.InitQueue(BUFFER_SIZE, NUM_WORKERS, taskRepo, urlRepo, taskResponseRepo)
	router := router.SetupRouter(db)
	fmt.Println("Starting server on port :8080")
	router.Run(":8080")
}
