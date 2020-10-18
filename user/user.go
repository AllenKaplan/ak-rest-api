package user

type UserService struct {
	Database string
}

type User struct {
	UserID int    `json:"UserID"`
	Name   string `json:"Name"`
	Email  string `json:"Email"`
}

type UsersResponse struct {
	Users []*User `json:"Users"`
}
