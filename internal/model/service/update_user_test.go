package service

import (
	"testing"

	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"github.com/rssistemasitu/crud-go/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_UpdateUserService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := mocks.NewMockUserRepository(controller)

	service := NewUserDomainService(repo)

	t.Run("when sending a valid user returns success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain.SetId(id)

		repo.EXPECT().UpdateUser(id, userDomain).Return(nil)

		err := service.UpdateUserService(id, userDomain)

		assert.Nil(t, err)
	})

	t.Run("when sending an invalid user returns error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain.SetId(id)

		repo.EXPECT().UpdateUser(id, userDomain).Return(
			rest_err.NewInternalServerError("error trying to update user"),
		)

		err := service.UpdateUserService(id, userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to update user")
	})

}
