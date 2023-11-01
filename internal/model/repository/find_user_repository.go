package repository

import (
	"context"
	"fmt"

	"github.com/rssistemasitu/crud-go/internal/configs/logger"
	"github.com/rssistemasitu/crud-go/internal/model"
	"github.com/rssistemasitu/crud-go/internal/model/repository/entity"
	"github.com/rssistemasitu/crud-go/internal/model/repository/entity/converter"
	"github.com/rssistemasitu/crud-go/internal/rest_err"
	"github.com/rssistemasitu/crud-go/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (
	model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserByEmail repository")

	collection_name := utils.GetEnvVariable(MONGODB_USERS_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err,
				zap.String("application", "user-application"),
				zap.String("flow", "find-user-by-email"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"

		logger.Error(errorMessage, err,
			zap.String("application", "user-application"),
			zap.String("flow", "find-user-by-email"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("application", "user-application"),
		zap.String("flow", "find-user-by-email"),
		zap.String("email", email))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(id string) (
	model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init findUserById repository")

	collection_name := utils.GetEnvVariable(MONGODB_USERS_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	object_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: object_id}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage, err,
				zap.String("application", "user-application"),
				zap.String("flow", "find-user-by-id"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"

		logger.Error(errorMessage, err,
			zap.String("application", "user-application"),
			zap.String("flow", "find-user-by-id"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserById repository executed successfully",
		zap.String("application", "user-application"),
		zap.String("flow", "find-user-by-id"),
		zap.String("userId", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}
