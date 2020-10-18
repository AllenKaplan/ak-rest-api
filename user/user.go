package user

var users []*User

type UserService struct {
	Database string
}

type User struct {
	ID    string `json:"userID"`
	Name  string `json:"name"`
	Email string `json:"name"`
}

type UsersResponse struct {
	Users []*User `json:"users"`
}
