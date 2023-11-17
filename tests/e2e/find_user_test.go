package e2e

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/controller"
	"github.com/rssistemasitu/crud-go/internal/model/repository"
	"github.com/rssistemasitu/crud-go/internal/model/service"
	"github.com/rssistemasitu/crud-go/tests/e2e/connection"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserController controller.UserControllerInterface
	Database       *mongo.Database
	Collection     *mongo.Collection
)

func TestMain(m *testing.M) {

	err := os.Setenv("MONGODB_USERS_COLLECTION", "test_user")
	if err != nil {
		return
	}

	closeConnection := func() {}

	Database, closeConnection = connection.OpenConnection()

	Collection = Database.Collection("test_user")

	repo := repository.NewUserRepository(Database)

	userService := service.NewUserDomainService(repo)

	UserController = controller.NewUserControllerInterface(userService)

	defer func() {
		os.Clearenv()
		closeConnection()
	}()

	code := m.Run()

	os.Exit(code)
}

func TestFindUserById(t *testing.T) {
	t.Run("when sending a valid id and controller returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)

		UserController.FindUserByIdController(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	})

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

		MakeRequest(ctx, param, url.Values{}, "GET", nil)

		UserController.FindUserByIdController(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})
}

func TestFindUserByEmail(t *testing.T) {
	t.Run("when sending a valid email and controller returns error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "testando_email@test.com",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)

		UserController.FindUserByEmailController(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	})

	t.Run("when sending a valid email and getting user controller returns success", func(t *testing.T) {
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
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)

		UserController.FindUserByEmailController(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)

	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser,
) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
