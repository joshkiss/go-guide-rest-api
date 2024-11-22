package main

import (
	"go-guide/rest-api/db"
	"go-guide/rest-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db.InitDB()
	server := gin.Default()

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("dotenv loaded")

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
