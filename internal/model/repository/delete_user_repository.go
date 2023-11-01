package repository

import (
	"context"

	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"github.com/rssistemasitu/crud-go/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {

	logger.Info("Init deleteUser repository")

	collection_name := utils.GetEnvVariable(MONGODB_USERS_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		logger.Error("Error trying to delete user",
			err,
			zap.String("application", "user-application"),
			zap.String("flow", "delete-user"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("DeleteUser repository executed successuly",
		zap.String("application", "user-application"),
		zap.String("flow", "delete-user"),
		zap.String("userId", userId))

	return nil
}
