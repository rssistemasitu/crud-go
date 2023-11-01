package request

// Para validar campos use a dependencia validator: go get github.com/go-playground/validator/v10
// Utilize binding
type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=10,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Age      int    `json:"age" binding:"required,min=1,max=150"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=4,max=100"`
	Age  int    `json:"age" binding:"omitempty,min=1,max=150"`
}
