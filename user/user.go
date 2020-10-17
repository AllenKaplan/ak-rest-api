package user

type UserService struct {
	Database string
}

type User struct {
	Name string `json:"name"`
}

type UsersResponse struct {
	Users []*User `json:"users"`
}
