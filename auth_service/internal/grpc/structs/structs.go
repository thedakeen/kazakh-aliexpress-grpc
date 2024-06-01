package structs

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=32"`
}

type RegisterRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=32"`
	Name     string `validate:"required,min=2,max=16"`
}

type IsAdminRequest struct {
	UserID string `validate:"required"`
}
