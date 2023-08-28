package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model/repository/entity"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init FindUserByEmail repository",
		zap.String("journey", "FindUserByEmail"))

	collection_name := os.Getenv(MONGO_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "FindUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		} else {
			errorMessage := "Error trying to find user by email"
			logger.Error(errorMessage,
				err,
				zap.String("journey", "FindUserByEmail"))
			return nil, rest_err.NewInternalServerError(errorMessage)
		}
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "FindUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertEntityToDomain(*userEntity), nil

}

func (ur *userRepository) FindUserByID(
	id string,
) (
	model.UserDomainInterface,
	*rest_err.RestErr,
) {
	logger.Info("Init FindUserByID repository",
		zap.String("journey", "FindUserByID"))

	collection_name := os.Getenv(MONGO_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "FindUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		} else {
			errorMessage := "Error trying to find user by id"
			logger.Error(errorMessage,
				err,
				zap.String("journey", "FindUserByID"))
			return nil, rest_err.NewInternalServerError(errorMessage)
		}
	}

	logger.Info("FindUserByID repository executed successfully",
		zap.String("journey", "FindUserByID"),
		zap.String("userId", userEntity.ID.Hex()),
	)

	return converter.ConvertEntityToDomain(*userEntity), nil

}
