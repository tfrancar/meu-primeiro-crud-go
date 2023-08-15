package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tfrancar/meu-primeiro-crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/create_User", controller.CreateUser)
	r.PUT("/update_User/:userId", controller.UpdateUser)
	r.DELETE("/delete_User/:userId", controller.DeleteUser)
}
