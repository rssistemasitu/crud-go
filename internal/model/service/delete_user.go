package service

import (
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"go.uber.org/zap"
)

func (*userDomainService) DeleteUser(string) *rest_err.RestErr {
	logger.Info("Delete user model",
		zap.String("application", "user-application"),
		zap.String("flow", "delete-create"))
	return nil
}
