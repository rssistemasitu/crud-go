package e2e

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/controller/model/request"
	"github.com/rssistemasitu/crud-go/internal/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateUser(t *testing.T) {
	t.Run("when sending a valid request returns success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		id := primitive.NewObjectID()
		email := "test@test.com"

		bm := bson.M{
			"_id":   id,
			"email": email,
			"name":  "Old_Name",
			"age":   10,
		}

		_, err := Collection.InsertOne(context.Background(), bm)
		if err != nil {
			t.Fatal(err)
		}

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}

		userRequest := request.UserUpdateRequest{
			Name: "Rogerio",
			Age:  25,
		}

		b, _ := json.Marshal(userRequest)

		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)

		UserController.UpdateUserController(ctx)
		assert.EqualValues(t, http.StatusOK, recorder.Result().StatusCode)

		userEntity := entity.UserEntity{}

		filter := bson.D{{Key: "_id", Value: id}}

		_ = Collection.FindOne(context.Background(), filter).Decode(&userEntity)

		assert.EqualValues(t, userEntity.Name, userRequest.Name)
		assert.EqualValues(t, userEntity.Age, userRequest.Age)

	})
}
