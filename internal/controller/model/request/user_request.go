package request

// Para validar campos use a dependencia validator: go get github.com/go-playground/validator/v10
// Utilize binding
type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,gte=4,lte=100"`
	Age      int    `json:"age" binding:"required,gte=1,lte=150"`
}
