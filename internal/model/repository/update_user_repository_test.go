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

func TestUserRepository_UpdateUser(t *testing.T) {
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

	mtestDb.Run("when sending a valid user should returns success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		domain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)
		domain.SetId(primitive.NewObjectID().Hex())

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		err := repo.UpdateUser(domain.GetId(), domain)

		assert.Nil(t, err)
	})

	mtestDb.Run("when sending a invalid user should returns error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		domain := model.NewUserDomain(
			"teste@test.com", "1234567*", "Test", 10,
		)
		domain.SetId(primitive.NewObjectID().Hex())

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		err := repo.UpdateUser(domain.GetId(), domain)

		assert.NotNil(t, err)
	})

}
