package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_DeleteUser(t *testing.T) {
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

	mtestDb.Run("when sending a valid user_id should returns success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		err := repo.DeleteUser("test")

		assert.Nil(t, err)
	})

	mtestDb.Run("when sending a invalid user_id should returns error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		database_mock := mt.Client.Database(database_name)

		repo := NewUserRepository(database_mock)

		err := repo.DeleteUser("test")

		assert.NotNil(t, err)
	})

}
