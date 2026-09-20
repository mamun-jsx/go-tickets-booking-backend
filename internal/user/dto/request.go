package dto

// by this we will transfer data
type CreateRequest struct {
	
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=12"`
}
