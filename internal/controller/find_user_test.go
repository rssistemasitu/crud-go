package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"github.com/rssistemasitu/crud-go/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)

	controller := NewUserControllerInterface(service)

	t.Run("when sending a valid email controller returns success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		ud := model.NewUserDomain(
			"test@test.com",
			"1234567*",
			"Test",
			10,
		)

		token, _ := ud.GenerateToken()

		context.Request.Header.Set("Authorization", token)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		service.EXPECT().FindUserByEmailService("test@test.com").Return(
			ud,
			nil,
		)

		MakeRequest(context, param, url.Values{}, "GET", nil)

		controller.FindUserByEmailController(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})

	t.Run("when sending an invalid email controller returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "TEST_ERROR",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)

		controller.FindUserByEmailController(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	})

	t.Run("when sending a valid email service returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		ud := model.NewUserDomain(
			"test@test.com",
			"1234567*",
			"Test",
			10,
		)

		token, _ := ud.GenerateToken()

		context.Request.Header.Set("Authorization", token)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: ud.GetEmail(),
			},
		}

		service.EXPECT().FindUserByEmailService(ud.GetEmail()).Return(
			nil,
			rest_err.NewInternalServerError("error test"),
		)

		MakeRequest(context, param, url.Values{}, "GET", nil)

		controller.FindUserByEmailController(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})
}

func TestUserControllerInterface_FindUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)

	controller := NewUserControllerInterface(service)

	t.Run("when sending a valid id controller returns success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userDomain := model.NewUserDomain(
			"test@test.com",
			"1234567*",
			"Test",
			10,
		)

		id := primitive.NewObjectID().Hex()

		userDomain.SetId(id)

		token, _ := userDomain.GenerateToken()

		context.Request.Header.Set("Authorization", token)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().FindUserByIdService(id).Return(
			userDomain,
			nil,
		)

		MakeRequest(context, param, url.Values{}, "GET", nil)

		controller.FindUserByIdController(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})

	t.Run("when sending an invalid id controller returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "TEST_ERROR",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)

		controller.FindUserByIdController(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when sending a valid id service returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userDomain := model.NewUserDomain(
			"test@test.com",
			"1234567*",
			"Test",
			10,
		)

		id := primitive.NewObjectID().Hex()

		userDomain.SetId(id)

		token, _ := userDomain.GenerateToken()

		context.Request.Header.Set("Authorization", token)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().FindUserByIdService(id).Return(
			nil,
			rest_err.NewInternalServerError("error test"),
		)

		MakeRequest(context, param, url.Values{}, "GET", nil)

		controller.FindUserByIdController(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)

	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser,
) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
