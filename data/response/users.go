package response

type AuthenticationResponse struct {
	Token string `json:"token" binding:"required"`
}

type GetUserResponse struct {
	ID        string `json:"id" binding:"omite" format:"uuid"`
	FirstName string `json:"firstName" binding:"omitempty,max=30"`
	LastName  string `json:"lastName" binding:"omitempty,max=30"`
	Username  string `json:"username" binding:"required,lowercase,alphanum,min=3,max=30"`
	Email     string `json:"email" binding:"required,email"`
}
