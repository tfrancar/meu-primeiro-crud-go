package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/tfrancar/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model"
	"github.com/tfrancar/meu-primeiro-crud-go/src/view"
	"go.uber.org/zap"
)

// Pode ser criada uma variável global ou um construtor para que o userdomain
var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUserServices(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "CreateUser"),
	)
	var userRequest request.UserResquest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "CreateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	// inicializando o dominio
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	// inicializando o service
	// Repassa o domain criado ali em cima para o CreateUser.
	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"),
		)
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "CreateUser"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	),
	)
}
