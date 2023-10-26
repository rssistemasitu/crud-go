package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/controller/routes"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error to initialize server", err)
	}
}
