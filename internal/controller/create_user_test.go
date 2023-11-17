package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/controller/model/request"
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"github.com/rssistemasitu/crud-go/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)

	controller := NewUserControllerInterface(service)

	t.Run("when sending a invalid userId controller returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{}

		userRequest := request.UserRequest{
			Email:    "TEST_ERROR",
			Password: "1234567*",
			Name:     "Test",
			Age:      10,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, param, url.Values{}, "POST", stringReader)

		controller.CreateUserController(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when sending a valid request service returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{}

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "1234567*",
			Name:     "Test",
			Age:      10,
		}

		domain := model.NewUserDomain(
			userRequest.Email,
			userRequest.Password,
			userRequest.Name,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().CreateUserService(domain).Return(
			nil,
			rest_err.NewInternalServerError("error test"),
		)

		MakeRequest(context, param, url.Values{}, "POST", stringReader)

		controller.CreateUserController(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("when sending a valid request service returns success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{}

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "1234567*",
			Name:     "Test",
			Age:      10,
		}

		domain := model.NewUserDomain(
			userRequest.Email,
			userRequest.Password,
			userRequest.Name,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().CreateUserService(domain).Return(
			domain,
			nil,
		)

		MakeRequest(context, param, url.Values{}, "POST", stringReader)

		controller.CreateUserController(context)

		assert.EqualValues(t, http.StatusCreated, recorder.Code)
	})
}
