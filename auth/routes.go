package auth

import (
	"errors"
	"fmt"
)

func (s *AuthService) Login(login *LoginRequest) (*Token, error) {
	//find user based on email
	retrievedLogin, err := s.Repo.getLogin(login.Email)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error retrieving login")
	}

	//validate snubmitted password is the retrievedLogin pw
	validLogin := retrievedLogin.Password == login.Password
	if !validLogin {
		return nil, errors.New("invalid login")
	}

	//check if token exists
	token, _ := s.Repo.retrieveToken(retrievedLogin.UserID)

	//if no token exists generate
	if token == nil {
		token, err = s.generateToken(retrievedLogin.UserID)
		if err != nil {
			return nil, fmt.Errorf("%v --> %s", err, "error generating token")
		}
	}

	return token, nil
}

func (s *AuthService) ValidateToken(token *Token) (bool, error) {
	//retrieve token from cache
	retrievedToken, err := s.Repo.retrieveToken(token.UserID)
	if err != nil {
		return false, fmt.Errorf("%v --> %s", err, "error retrieving token")
	}

	//compare sentToken and retrievedToken
	if token.TokenID == retrievedToken.TokenID {
		return true, nil
	}

	return false, fmt.Errorf("%s", "could not validate token")
}

func (s *AuthService) Create(login *Login) (*Token, error) {
	_, err := s.Repo.createLogin(login)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error creating login")
	}

	token, err := s.generateToken(login.UserID)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error generating token")
	}

	return token, nil
}

func (s *AuthService) generateToken(userID int) (*Token, error) {
	generatedToken := generateToken(userID)

	_, err := s.Repo.storeToken(generatedToken)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error storing token")
	}
	return generatedToken, nil
}
