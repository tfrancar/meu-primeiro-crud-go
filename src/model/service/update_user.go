package service

import (
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUserService model", zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUserServices"))
		return err
	}

	logger.Info("updateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	return nil
}
