package service

import (
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIdService(
	userId string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserById services",
		zap.String("application", "user-application"),
		zap.String("flow", "find-user-by-idservices"))
	return ud.userRepository.FindUserById(userId)
}

func (ud *userDomainService) FindUserByEmailService(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail services",
		zap.String("application", "user-application"),
		zap.String("flow", "find-user-by-email-services"))
	return ud.userRepository.FindUserByEmail(email)
}
