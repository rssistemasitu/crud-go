package service

import (
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"go.uber.org/zap"
)

func (*userDomainService) FindUser(userId string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Find user model",
		zap.String("application", "user-application"),
		zap.String("flow", "find-user"))
	return nil, nil
}
