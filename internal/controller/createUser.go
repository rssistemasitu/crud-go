package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/controller/model/request"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return

	}
	fmt.Print(userRequest)
}
