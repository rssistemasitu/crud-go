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

func TestUserDomainService_CreateUserService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := mocks.NewMockUserRepository(controller)

	service := NewUserDomainService(repo)

	t.Run("when not exists user returns success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain.SetId(id)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(
			nil,
			nil)

		repo.EXPECT().CreateUser(userDomain).Return(
			userDomain,
			nil)

		user, err := service.CreateUserService(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, user.GetId(), userDomain.GetId())
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
	})

	t.Run("when exists user returns error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain.SetId(id)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateUserService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email is already registered in another account.")
	})

	t.Run("when not exists user returns error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain.SetId(id)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(
			nil,
			nil)

		repo.EXPECT().CreateUser(userDomain).Return(
			nil,
			rest_err.NewInternalServerError("Error trying to create user."))

		user, err := service.CreateUserService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to create user.")
	})

}
