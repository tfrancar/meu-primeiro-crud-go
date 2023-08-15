package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Voce chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
