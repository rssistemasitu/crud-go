package main

import (
	"github.com/rssistemasitu/crud-go/internal/controller"
	"github.com/rssistemasitu/crud-go/internal/model/repository"
	"github.com/rssistemasitu/crud-go/internal/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
