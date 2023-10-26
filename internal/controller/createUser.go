package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/configs/validation"
	"github.com/rssistemasitu/crud-go/internal/controller/model/request"
	"github.com/rssistemasitu/crud-go/internal/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("application", "user-application"), zap.String("event", "user-create-controller"))
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user", err)
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	response := response.UserResponse{
		ID:    "test-id",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	c.JSON(http.StatusCreated, response)
}
