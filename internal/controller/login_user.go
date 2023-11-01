package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/configs/validation"
	"github.com/rssistemasitu/crud-go/internal/controller/model/request"
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUserController(c *gin.Context) {
	logger.Info("Init LoginUser controller",
		zap.String("application", "user-application"),
		zap.String("flow", "user-login-controller"))

	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, err := uc.service.LoginUserService(domain)

	if err != nil {
		logger.Error("Error trying to call loginUser service",
			err,
			zap.String("application", "user-application"),
			zap.String("flow", "user-login-controller"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("LoginUser controller executed successfully",
		zap.String("application", "user-application"),
		zap.String("flow", "user-login-email"),
		zap.String("userId", domainResult.GetId()))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
