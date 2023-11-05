package repository

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/rssistemasitu/crud-go/internal/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserById(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	err := os.Setenv("MONGODB_USERS_COLLECTION", collection_name)

	if err != nil {
		t.FailNow()
		return
	}

	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when sending valid user_id should returns success", func(mt *mtest.T) {

		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@test.com",
			Password: "1234567*",
			Name:     "Test",
			Age:      10,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		userFound, _ := repo.FindUserById(userEntity.ID.Hex())

		assert.NotNil(t, userFound)
		assert.EqualValues(t, userFound.GetId(), userEntity.ID.Hex())
		assert.EqualValues(t, userFound.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userFound.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userFound.GetName(), userEntity.Name)
		assert.EqualValues(t, userFound.GetAge(), userEntity.Age)
	})

	mtestDb.Run("when sending some user_id mongodb should returns error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		userFound, err := repo.FindUserById("123456879")

		assert.NotNil(t, err)
		assert.Nil(t, userFound)
		assert.Equal(t, err.Code, http.StatusInternalServerError)
		assert.Equal(t, err.Message, "Error trying to find user by id")
	})

	mtestDb.Run("when sending some user_id and mongodb should returns error No Documents", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		userFound, err := repo.FindUserById("123456879")

		assert.NotNil(t, err)
		assert.Nil(t, userFound)
		assert.Equal(t, err.Code, http.StatusNotFound)
		assert.Equal(t, err.Message, fmt.Sprintf("User not found with this id: %s", "123456879"))
	})
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	err := os.Setenv("MONGODB_USERS_COLLECTION", collection_name)

	if err != nil {
		t.FailNow()
		return
	}

	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDb.Close()

	mtestDb.Run("when sending valid email should returns success", func(mt *mtest.T) {

		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@test.com",
			Password: "1234567*",
			Name:     "Test",
			Age:      10,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		userFound, _ := repo.FindUserByEmail(userEntity.Email)

		assert.NotNil(t, userFound)
		assert.EqualValues(t, userFound.GetId(), userEntity.ID.Hex())
		assert.EqualValues(t, userFound.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userFound.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userFound.GetName(), userEntity.Name)
		assert.EqualValues(t, userFound.GetAge(), userEntity.Age)
	})

	mtestDb.Run("when sending some email mongodb should returns error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		userFound, err := repo.FindUserByEmail("some_email@test.com")

		assert.NotNil(t, err)
		assert.Nil(t, userFound)
		assert.Equal(t, err.Code, http.StatusInternalServerError)
		assert.Equal(t, err.Message, "Error trying to find user by email")
	})

	mtestDb.Run("when sending some email and mongodb should returns error No Documents", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", database_name, collection_name),
			mtest.FirstBatch))

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		userFound, err := repo.FindUserByEmail("some_email@test.com")

		assert.NotNil(t, err)
		assert.Nil(t, userFound)
		assert.Equal(t, err.Code, http.StatusNotFound)
		assert.Equal(t, err.Message, fmt.Sprintf("User not found with this email: %s", "some_email@test.com"))
	})
}

func convertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
