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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_LoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)

	controller := NewUserControllerInterface(service)

	t.Run("when sending a invalid request login controller returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{}

		userLogin := request.UserLogin{
			Email:    "TEST_ERROR",
			Password: "1234",
		}

		b, _ := json.Marshal(userLogin)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, param, url.Values{}, "POST", stringReader)

		controller.LoginUserController(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when sending a valid request login and service returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{}

		userLogin := request.UserLogin{
			Email:    "test@test.com",
			Password: "1234567*",
		}

		b, _ := json.Marshal(userLogin)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		domain := model.NewUserLoginDomain(
			userLogin.Email,
			userLogin.Password,
		)

		service.EXPECT().LoginUserService(domain).Return(
			nil,
			"",
			rest_err.NewInternalServerError("error test"),
		)

		MakeRequest(context, param, url.Values{}, "POST", stringReader)

		controller.LoginUserController(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("when sending a valid request login and controller returns success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{}

		userLogin := request.UserLogin{
			Email:    "test@test.com",
			Password: "1234567*",
		}

		b, _ := json.Marshal(userLogin)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		domain := model.NewUserLoginDomain(
			userLogin.Email,
			userLogin.Password,
		)

		domainResult := model.NewUserLoginDomain(
			userLogin.Email,
			userLogin.Password,
		)
		domainResult.SetId(id)

		service.EXPECT().LoginUserService(domain).Return(
			domainResult,
			id,
			nil,
		)

		MakeRequest(context, param, url.Values{}, "POST", stringReader)

		controller.LoginUserController(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, recorder.Header().Values("Authorization")[0], id)
	})
}
