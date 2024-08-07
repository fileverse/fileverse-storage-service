package main

import (
	"fmt"
	"storage/src/config"
	"storage/src/pkg/db"
	"storage/src/pkg/logger"
	"storage/src/services/server"

	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Initializing server")
	godotenv.Load()
	config.Setup("./src/config")

	logger.Setup()
	db.Setup()
	fmt.Println("Server initialized")
}

func main() {
	fmt.Println("Starting server")
	s := server.Build()
	s.Run(":8080")
}
