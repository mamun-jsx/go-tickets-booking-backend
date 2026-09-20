package dto

// send the response to the client but not password
type Response struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CreatedAt string `json:"created_at"`
	
}