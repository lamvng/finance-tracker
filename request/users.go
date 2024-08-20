package request

type CreateUserInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Username  string `json:"userName" binding:"required,lowercase,alphanum"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type AuthenticationInput struct {
	Username string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}
