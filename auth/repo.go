package auth

import (
	"golang.org/x/crypto/bcrypt"
)

var logins map[int]*Login
var tokens map[int]*Token

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s AuthService) getLogin(userID int) (*Login, error) {
	return logins[userID], nil
}

func (s AuthService) createLogin(login *Login) (*Login, error) {
	if logins == nil {
		logins = make(map[int]*Login)
	}
	logins[login.UserID] = login
	return login, nil
}

func (s AuthService) storeToken(token *Token) (*Token, error) {
	if tokens == nil {
		tokens = make(map[int]*Token)
	}
	tokens[token.UserID] = token

	return token, nil
}

func (s AuthService) retrieveToken(userID int) (*Token, error) {
	return tokens[userID], nil
}
