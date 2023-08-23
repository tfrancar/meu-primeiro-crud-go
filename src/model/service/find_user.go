package service

import (
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface,
	*rest_err.RestErr,
) {
	return nil, nil
}
