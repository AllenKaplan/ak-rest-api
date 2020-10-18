package auth

type AuthService struct {
	Database string
}

type Token struct {
	userID  int
	tokenID int
	expiry  int
}

type Login struct {
	userID   int    `json:"userID"`
	Email    string `json:"name"`
	Password string `json:"name"`
}
