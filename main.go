package main

import (
	"fmt"
	"github.com/gitnoober/grawler/router"
)

func main() {
	router := router.SetupRouter()
	fmt.Println("Starting server on port :8080")
	router.Run(":8080")
}
