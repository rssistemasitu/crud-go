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

func TestUserDomainService_FindUserByIdService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := mocks.NewMockUserRepository(controller)

	service := NewUserDomainService(repo)

	t.Run("when exists user returns success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain.SetId(id)

		repo.EXPECT().FindUserById(id).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByIdService(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetId(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when not exists user returns error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repo.EXPECT().FindUserById(id).Return(nil, rest_err.NewNotFoundError("user not found"))

		userDomainReturn, err := service.FindUserByIdService(id)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := mocks.NewMockUserRepository(controller)

	service := NewUserDomainService(repo)

	t.Run("when exists user returns success", func(t *testing.T) {
		id := "teste@test.com"
		email := "teste@test.com"

		userDomain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain.SetId(id)

		repo.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByEmailService(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetId(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
		assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())
	})

	t.Run("when not exists user returns error", func(t *testing.T) {
		email := "teste@test.com"

		repo.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))

		userDomainReturn, err := service.FindUserByEmailService(email)

		assert.Nil(t, userDomainReturn)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "user not found")
	})
}
