package dto

// by this we will transfer data
type CreateRequest struct {
	// register | create user 
	
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=12"`
}



// login user 
type LoginRequest struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=12"`
}
