package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"github.com/rssistemasitu/crud-go/internal/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// Usando UUID - go get github.com/google/uuid

func (uc *userControllerInterface) FindUserByIdController(c *gin.Context) {
	logger.Info("Init findUserById controller",
		zap.String("application", "user-application"),
		zap.String("flow", "find-user-by-id"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId", err,
			zap.String("application", "user-application"),
			zap.String("flow", "find-user-by-id"))

		errorMessage := rest_err.NewBadRequestError("UserId is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdService(userId)
	if err != nil {
		logger.Error("Error trying to call findUserById service", err,
			zap.String("application", "user-application"),
			zap.String("flow", "find-user-by-id"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserById controller executed successfully",
		zap.String("application", "user-application"),
		zap.String("flow", "find-user-by-id"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmailController(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("application", "user-application"),
		zap.String("flow", "find-user-by-email"))

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate e-mail", err,
			zap.String("application", "user-application"),
			zap.String("flow", "find-user-by-email"))

		errorMessage := rest_err.NewBadRequestError("UserEmail is not a valid e-mail")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmail service", err,
			zap.String("application", "user-application"),
			zap.String("flow", "find-user-by-email"))

		c.JSON(err.Code, err)
		return
	}

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("application", "user-application"),
		zap.String("flow", "find-user-by-email"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))

}
