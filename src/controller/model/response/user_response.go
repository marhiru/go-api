package response

type UserResponse struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
	Age   int8   `json:"age,omitempty"`
}
