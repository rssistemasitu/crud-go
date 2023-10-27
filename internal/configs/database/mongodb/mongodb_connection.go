package mongodb

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL  = "MONGODB_URL"
	MONGODB_NAME = "MONGODB_NAME"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_name := getEnvVariable(MONGODB_NAME)
	mongodb_uri := getEnvVariable(MONGODB_URL)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Init connection with Mongo database")
	return client.Database(mongodb_name), nil

}

func getEnvVariable(logName string) string {
	err := godotenv.Load("configs/.env")
	if err != nil {
		logger.Error("Error loading .env file", err)
	}
	return os.Getenv(logName)
}
