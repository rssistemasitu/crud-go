package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/controller"
	"github.com/rssistemasitu/crud-go/internal/controller/routes"
	"github.com/rssistemasitu/crud-go/internal/model/service"
)

func main() {
	// inicia as dependencias
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error to initialize server", err)
	}
}
