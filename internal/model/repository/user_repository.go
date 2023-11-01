package repository

import (
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USERS_COLLECTION = "MONGODB_USERS_COLLECTION"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserById(id string) (
		model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmail(email string) (
		model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr

	DeleteUser(id string) *rest_err.RestErr
}
