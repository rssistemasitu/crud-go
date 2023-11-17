package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"github.com/rssistemasitu/crud-go/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)

	controller := NewUserControllerInterface(service)

	t.Run("when sending a valid id controller returns success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().DeleteUserService(id).Return(
			nil,
		)

		MakeRequest(context, param, url.Values{}, "DELETE", nil)

		controller.DeleteUserController(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

	t.Run("when sending a invalid userId controller returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "teste",
			},
		}

		MakeRequest(context, param, url.Values{}, "DELETE", nil)

		controller.DeleteUserController(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("when sending a valid id service returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().DeleteUserService(id).Return(
			rest_err.NewInternalServerError("error test"),
		)

		MakeRequest(context, param, url.Values{}, "DELETE", nil)

		controller.DeleteUserController(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})
}
