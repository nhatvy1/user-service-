package vo

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6,max=64"`
	FirstName string `json:"firstname" binding:"required,min=1,max=50"`
	LastName  string `json:"lastname" binding:"required,min=1,max=50"`
}
