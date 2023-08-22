package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/tfrancar/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

// Pode ser criada uma variável global ou um construtor para que o userdomain
var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
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

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"))

	c.String(http.StatusOK, "")
}
