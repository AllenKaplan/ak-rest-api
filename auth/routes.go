package auth

import (
	"errors"
	"fmt"
	"math/rand"
)

func (s AuthService) Login(login *Login) (*Token, error) {
	retrievedLogin, err := s.getLogin(login.UserID)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error retrieving login")
	}

	validLogin := retrievedLogin.Password == login.Password

	if !validLogin {
		return nil, errors.New("invalid login")
	}

	token, err := s.generateToken(login.UserID)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error generating token")
	}

	return token, nil
}

func (s AuthService) ValidateToken(token *Token) (bool, error) {
	retrievedToken, err := s.retrieveToken(token.UserID)
	if err != nil {
		return false, fmt.Errorf("%v --> %s", err, "error retrieving token")
	}

	if token.TokenID == retrievedToken.TokenID {
		return true, nil
	}

	return false, fmt.Errorf("%s", "could not validate token")
}

func (s AuthService) Create(login *Login) (*Token, error) {
	_, err := s.createLogin(login)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error creating login")
	}

	token, err := s.generateToken(login.UserID)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error generating token")
	}

	return token, nil
}

func (s AuthService) generateToken(userID int) (*Token, error) {
	generatedToken := &Token{
		TokenID: rand.Intn(100000) + 1000000,
		UserID:  userID,
		Expiry:  0,
	}

	_, err := s.storeToken(generatedToken)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error storing token")
	}
	return generatedToken, nil
}
