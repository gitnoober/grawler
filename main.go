package main

import (
	"fmt"
	"os"

	"github.com/gitnoober/grawler/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	router := router.SetupRouter(db)
	fmt.Println("Starting server on port :8080")
	router.Run(":8080")
}
