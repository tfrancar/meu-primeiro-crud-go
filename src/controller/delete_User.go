package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {

	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("userId")
	if err := c.ShouldBindJSON(&userRequest); err != nil ||
		strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	// inicializando o dominio
	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)
	// inicializando o service
	// Repassa o domain criado ali em cima para o updateUserService.
	err := uc.service.UpdateUserServices(userId, domain)
	if err != nil {
		logger.Error(
			"Error trying to call updateUser service",
			err,
			zap.String("journey", "updateUser"),
		)
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)

}
