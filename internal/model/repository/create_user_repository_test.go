package repository

import (
	"os"
	"testing"

	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {
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

	mtestDb.Run("when sending a valid domain should returns success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		domain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain, err := repo.CreateUser(domain)

		_, errId := primitive.ObjectIDFromHex(userDomain.GetId())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), domain.GetEmail())
		assert.EqualValues(t, userDomain.GetPassword(), domain.GetPassword())
		assert.EqualValues(t, userDomain.GetName(), domain.GetName())
		assert.EqualValues(t, userDomain.GetAge(), domain.GetAge())
	})

	mtestDb.Run("when sending a invalid domain should returns error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		domain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)

		userDomain, err := repo.CreateUser(domain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

}
