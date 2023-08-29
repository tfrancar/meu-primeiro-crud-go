package service

import (
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository //dependencia
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDServices(id string) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailServices(email string) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUserServices(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}
