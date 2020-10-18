package auth

type AuthService struct {
	Database string
}

type Token struct {
	UserID  int
	TokenID int
	Expiry  int
}

type Login struct {
	UserID   int    `json:"userID"`
	Email    string `json:"email"`
	Password string `json:"name"`
}
