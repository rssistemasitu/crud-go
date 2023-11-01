package service

import (
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/model/repository"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (
		model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIdService(string) (
		model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailService(string) (
		model.UserDomainInterface, *rest_err.RestErr)

	UpdateUserService(string, model.UserDomainInterface) *rest_err.RestErr

	DeleteUserService(string) *rest_err.RestErr

	LoginUserService(model.UserDomainInterface) (
		model.UserDomainInterface, *rest_err.RestErr)
}
