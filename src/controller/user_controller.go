package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model/service"
)

// Metódo para iniciar o controller
func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

// contém os metódos do nosso controller.
type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
