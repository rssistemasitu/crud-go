package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/model/service"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByIdController(c *gin.Context)
	FindUserByEmailController(c *gin.Context)
	CreateUserController(c *gin.Context)
	UpdateUserController(c *gin.Context)
	DeleteUserController(c *gin.Context)
	LoginUserController(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
