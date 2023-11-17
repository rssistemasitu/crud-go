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

func TestUserDomainService_LoginUserService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := mocks.NewMockUserRepository(controller)

	service := NewUserDomainService(repo)

	t.Run("when not exists email returns error", func(t *testing.T) {

		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain.SetId(id)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(
			nil,
			rest_err.NewForbiddenError("Unauthorized"))

		user, token, err := service.LoginUserService(userDomain)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.EqualValues(t, err.Message, "Unauthorized")
		assert.Empty(t, token)
	})

	t.Run("when not check password returns error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain.SetId(id)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(
			userDomain,
			nil)

		user, token, err := service.LoginUserService(userDomain)

		assert.NotNil(t, err)
		assert.Nil(t, user)
		assert.EqualValues(t, err.Message, "Unauthorized")
		assert.Empty(t, token)
	})

	t.Run("when generate token returns error", func(t *testing.T) {
		userDomainMock := mocks.NewMockUserDomainInterface(controller)

		userDomainMock.EXPECT().GetEmail().Return("teste@test.com")
		userDomainMock.EXPECT().GetPassword().Return("1234567*")
		// userDomainMock.EXPECT().EncryptPassword()
		userDomainMock.EXPECT().CheckPasswordHash("1234567*").Return(true)

		userDomainMock.EXPECT().GenerateToken().Return(
			"",
			rest_err.NewInternalServerError("Error trying to generate jwt token"),
		)

		repo.EXPECT().FindUserByEmail("teste@test.com").Return(
			userDomainMock,
			nil)

		user, token, err := service.LoginUserService(userDomainMock)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to generate jwt token")

	})

	t.Run("when sending an valid email and password returns success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain.SetId(id)

		repo.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(
			userDomain,
			nil)

		userDomainMock := model.NewUserDomain(
			userDomain.GetEmail(),
			userDomain.GetPassword(),
			userDomain.GetName(),
			userDomain.GetAge(),
		)

		userDomain.EncryptPassword()
		userDomainMock.CheckPasswordHash(userDomain.GetPassword())

		user, token, err := service.LoginUserService(userDomainMock)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.NotEmpty(t, token)
	})

}
