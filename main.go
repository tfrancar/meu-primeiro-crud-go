package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/database/mongodb"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/tfrancar/meu-primeiro-crud-go/src/controller/routes"
)

func main() {
	gin.SetMode(gin.DebugMode)
	logger.Info("About to start application")
	err := godotenv.Load() //Variavel que vai armazenar o erro
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error connection", err)
	}
}
