package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tfrancar/meu-primeiro-crud-go/src/controller"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {

	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/create_User", userController.CreateUser)
	r.PUT("/update_User/:userId", userController.UpdateUser)
	r.DELETE("/delete_User/:userId", userController.DeleteUser)
}
