package request

type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      int8   `json:"age"`
}
