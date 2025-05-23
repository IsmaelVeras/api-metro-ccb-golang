package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email,max=60"`
	Password string `json:"password" binding:"required,min=4,max=40,containsany=!@#$&"`
	Name     string `json:"name" binding:"required,min=2,max=20"`
	Age      int8   `json:"age" binding:"required,min=1,max=120"`
}
