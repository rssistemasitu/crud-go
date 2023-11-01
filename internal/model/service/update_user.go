package service

import (
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser service",
		zap.String("application", "user-application"),
		zap.String("flow", "update-user"))

	err := ud.userRepository.UpdateUser(userId, userDomain)

	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("application", "user-application"),
			zap.String("flow", "update-user"))
		return err
	}

	logger.Info("UpdateUser service executed successfully",
		zap.String("application", "user-application"),
		zap.String("flow", "update-user"))

	return nil
}
