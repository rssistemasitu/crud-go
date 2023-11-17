package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUserController(c *gin.Context) {
	logger.Info("Init DeleteUser controller",
		zap.String("application", "user-application"),
		zap.String("flow", "user-delete-controller"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be hex")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := uc.service.DeleteUserService(userId)

	if err != nil {
		logger.Error("Error trying to call deleteUser service",
			err,
			zap.String("application", "user-application"),
			zap.String("flow", "user-delete-controller"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfuly",
		zap.String("application", "user-application"),
		zap.String("flow", "user-delete-controller"),
		zap.String("userId", userId))

	c.Status(http.StatusOK)
}
