package service

import (
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init LoginUser service",
		zap.String("application", "user-application"),
		zap.String("flow", "user-login-service"))

	user, err := ud.FindUserByEmailService(userDomain.GetEmail())

	if err != nil {
		return nil, "", rest_err.NewForbiddenError("Unauthorized")
	}

	if !user.CheckPasswordHash(userDomain.GetPassword()) {
		return nil, "", rest_err.NewForbiddenError("Unauthorized")
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info("LoginUser service executed successfully",
		zap.String("application", "user-application"),
		zap.String("flow", "user-login-service"),
		zap.String("userId", user.GetId()))

	return user, token, nil
}
