package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/tfrancar/meu-primeiro-crud-go/src/controller/routes"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	logger.Info("About to start application")
	err := godotenv.Load() //Variavel que vai armazenar o erro
	if err != nil {
		logger.Error("Error loading .env file", err)
	}
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error connection", err)
	}
}
