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
)

func TestCreateUser(t *testing.T) {
	t.Run("when sending a valid email that already exists returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		email := "test@test.com"

		bm := bson.M{
			"email": email,
			"name":  t.Name(),
		}

		_, err := Collection.InsertOne(context.Background(), bm)
		if err != nil {
			t.Fatal(err)
		}

		userRequest := request.UserRequest{
			Email:    email,
			Password: "1234567*",
			Name:     "Test",
			Age:      10,
		}

		b, _ := json.Marshal(userRequest)

		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)

		UserController.CreateUserController(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	})

	t.Run("when sending a valid email that not exists returns success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		email := "testando@test.com"

		userRequest := request.UserRequest{
			Email:    email,
			Password: "1234567*",
			Name:     "Test",
			Age:      10,
		}

		b, _ := json.Marshal(userRequest)

		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)

		UserController.CreateUserController(ctx)

		userEntity := entity.UserEntity{}

		filter := bson.D{{Key: "email", Value: email}}

		_ = Collection.FindOne(context.Background(), filter).Decode(&userEntity)

		assert.EqualValues(t, http.StatusCreated, recorder.Code)
		assert.NotNil(t, userEntity.ID)
		assert.EqualValues(t, userEntity.Email, userRequest.Email)
		assert.EqualValues(t, userEntity.Name, userRequest.Name)
		assert.EqualValues(t, userEntity.Age, userRequest.Age)
	})
}
