package request

type AuthenticationRequest struct {
	Username string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	FirstName string `json:"firstName" binding:"required,max=30"`
	LastName  string `json:"lastName" binding:"required,max=30"`
	Username  string `json:"userName" binding:"required,lowercase,alphanum,min=3,max=30"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=30"`
}

type UpdateUserRequest struct {
	ID        string `json:"id" binding:"required" format:"uuid"`
	FirstName string `json:"firstName" binding:"omitempty,max=30"`
	LastName  string `json:"lastName" binding:"omitempty,max=30"`
	Username  string `json:"userName" binding:"omitempty,lowercase,alphanum,min=3,max=30"`
	Email     string `json:"email" binding:"omitempty,email"`
	Password  string `json:"password" binding:"omitempty,min=8,max=30"`
}
