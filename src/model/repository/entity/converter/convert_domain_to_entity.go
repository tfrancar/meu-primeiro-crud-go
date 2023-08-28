package converter

import (
	"github.com/tfrancar/meu-primeiro-crud-go/src/model"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}

}
