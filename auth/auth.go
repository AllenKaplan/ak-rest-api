package auth

type AuthService struct {
	Repo *AuthRepo
}

func NewService() *AuthService {
	return &AuthService{
		Repo: NewAuthRepo(),
	}
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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
