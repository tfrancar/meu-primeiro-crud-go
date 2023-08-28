package main

import (
	"github.com/tfrancar/meu-primeiro-crud-go/src/controller"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model/repository"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	//Init dependecies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)
	// return controller.NewUserControllerInterface(service), nil --> pode ser desta forma também.
	return userController
}
