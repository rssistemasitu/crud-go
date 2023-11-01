package service

import (
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(userId string) *rest_err.RestErr {
	logger.Info("Init Delete user model",
		zap.String("application", "user-application"),
		zap.String("flow", "delete-user"))

	err := ud.userRepository.DeleteUser(userId)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("application", "user-application"),
			zap.String("flow", "delete-user"))
		return err
	}

	return nil
}
