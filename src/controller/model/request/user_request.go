package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email,min=6,max=100"`
	Name     string `json:"name" binding:"required,min=4,max=40"`
	Password string `json:"password" binding:"required,containsany=!@#$*,min=6,max=16"`
	Age      int8   `json:"age" binding:"required,min=18,max=100"`
}
