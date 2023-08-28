package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/database/mongodb"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/tfrancar/meu-primeiro-crud-go/src/controller"
	"github.com/tfrancar/meu-primeiro-crud-go/src/controller/routes"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model/repository"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model/service"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	logger.Info("About to start application")
	err := godotenv.Load() //Variavel que vai armazenar o erro
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
	}

	//Init dependecies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8081"); err != nil {
		logger.Error("Error connection", err)
	}
}
