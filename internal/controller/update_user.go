package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/configs/validation"
	"github.com/rssistemasitu/crud-go/internal/controller/model/request"
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUserController(c *gin.Context) {
	logger.Info("Init UpdateUser controller",
		zap.String("application", "user-application"),
		zap.String("flow", "user-update-controller"))

	var userUpdateRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		logger.Error("Error trying to validate user", err)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be hex")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userUpdateRequest.Name,
		userUpdateRequest.Age,
	)

	err := uc.service.UpdateUserService(userId, domain)

	if err != nil {
		logger.Error("Error trying to call updateUser service",
			err,
			zap.String("application", "user-application"),
			zap.String("flow", "user-update-controller"))

		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed successfuly",
		zap.String("application", "user-application"),
		zap.String("flow", "user-update-controller"),
		zap.String("userId", userId))

	c.Status(http.StatusOK)
}
