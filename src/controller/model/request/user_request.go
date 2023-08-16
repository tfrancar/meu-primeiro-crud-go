package request

type UserResquest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,min=4,max=10"`
	Age      int8   `json:"age" binding:"required,min=1,max=200"`
}
