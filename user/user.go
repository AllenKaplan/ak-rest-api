package user

type UserService struct {
	Database string
}

type User struct {
	UserID int    `json:"userID"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UsersResponse struct {
	Users []*User `json:"users"`
}
