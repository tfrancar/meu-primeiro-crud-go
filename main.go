package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() //Variavel que vai armazenar o erro
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	routes := gin.Default()
	routes.InitRoutes()
}
