package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/getUserById/:userId", userController.FindUserByIdController)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmailController)
	r.POST("/createUser", userController.CreateUserController)
	r.PUT("/updateUser/:userId", userController.UpdateUserController)
	r.DELETE("/deleteUser/:userId", userController.DeleteUserController)
	r.POST("/login", userController.LoginUserController)
}
