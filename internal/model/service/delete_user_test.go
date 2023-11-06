package service

import (
	"testing"

	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"github.com/rssistemasitu/crud-go/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_DeleteUserService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := mocks.NewMockUserRepository(controller)

	service := NewUserDomainService(repo)

	t.Run("when sending a valid user_id returns success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repo.EXPECT().DeleteUser(id).Return(nil)

		err := service.DeleteUserService(id)

		assert.Nil(t, err)
	})

	t.Run("when sending an invalid user_id returns error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repo.EXPECT().DeleteUser(id).Return(
			rest_err.NewInternalServerError("Error trying to call repository"),
		)

		err := service.DeleteUserService(id)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to call repository")
	})
}
