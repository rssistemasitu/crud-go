package e2e

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteUser(t *testing.T) {
	t.Run("when sending a valid id and getting user controller returns success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID()

		b := bson.M{
			"_id":   id,
			"email": "test@test.com",
			"name":  t.Name(),
		}

		_, err := Collection.InsertOne(context.Background(), b)
		if err != nil {
			t.Fatal(err)
		}

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}

		MakeRequest(ctx, param, url.Values{}, "DELETE", nil)

		UserController.DeleteUserController(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

		filter := bson.D{{Key: "_id", Value: id}}
		result := Collection.FindOne(context.Background(), filter)

		assert.NotNil(t, result.Err())

	})
}
