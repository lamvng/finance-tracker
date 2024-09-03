package response

type GetUserByIDResponse struct {
	ID        string `json:"id" binding:"required" format:"uuid"`
	FirstName string `json:"firstName" binding:"required,max=30"`
	LastName  string `json:"lastName" binding:"required,max=30"`
	Username  string `json:"userName" binding:"required,lowercase,alphanum,min=3,max=30"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=30"`
}
