package request

// Pacote request com o modelo do usuário que será gravado no banco
// binding é uma requisição do gingonic
// validator requer que seja utilizado com o validate
type UserResquest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,min=4,max=10"`
	Age      int8   `json:"age" binding:"required,min=1,max=200"`
}

// Pacote request com o modelo do usuário que será gravado no banco
// binding é uma requisição do gingonic
// validator requer que seja utilizado com o validate
type UserUpdateResquest struct {
	Name string `json:"name" binding:"omitempty,min=4,max=10"`
	Age  int8   `json:"age" binding:"omitempty,min=1,max=200"`
}
