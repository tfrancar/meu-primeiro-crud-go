package converter

import (
	"github.com/tfrancar/meu-primeiro-crud-go/src/model"
	"github.com/tfrancar/meu-primeiro-crud-go/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex()) // retorna uma string, porém apenas o que interessa.
	return domain
}
