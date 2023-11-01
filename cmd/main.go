package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/configs/database/mongodb"
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/controller"
	"github.com/rssistemasitu/crud-go/internal/controller/routes"
	"github.com/rssistemasitu/crud-go/internal/model/repository"
	"github.com/rssistemasitu/crud-go/internal/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

// Atualizando configurações sem reiniciar a aplicação: go get https://github.com/fsnotify/fsnotify

func main() {
	// inicia as dependencias
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.Error("Error trying to connect to database", err)
		return
	}

	userController := initDependencies(database)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error to initialize server", err)
	}
}

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
