package auth

import (
	"errors"
)

type AuthRepo struct {
	LoginRepo map[int]*Login
	TokenRepo map[int]*Token
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{
		LoginRepo: make(map[int]*Login),
		TokenRepo: make(map[int]*Token),
	}
}

//database actions

func (r *AuthRepo) getLogin(email string) (*Login, error) {
	for userID, login := range r.LoginRepo {
		if login.Email == email {
			return r.LoginRepo[userID], nil
		}
	}

	return nil, errors.New("Could not find or retirieve user of given email")
}

func (r *AuthRepo) createLogin(login *Login) (*Login, error) {
	// fmt.Println("Login being created | ", login.Email)
	r.LoginRepo[login.UserID] = login
	return login, nil
}

func (r *AuthRepo) storeToken(token *Token) (*Token, error) {
	r.TokenRepo[token.UserID] = token

	return token, nil
}

func (r *AuthRepo) retrieveToken(userID int) (*Token, error) {
	resultToken, ok := r.TokenRepo[userID]
	if !ok {
		return nil, errors.New("Could not find or retireve token")
	}
	return resultToken, nil
}
