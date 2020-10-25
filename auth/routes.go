package auth

import (
	"errors"
	"fmt"
)

func (s *AuthService) Login(login *LoginRequest) (string, error) {
	//find user based on email
	retrievedLogin, err := s.Repo.getLogin(login.Email)
	if err != nil {
		return "", fmt.Errorf("%v --> %s", err, "error retrieving login")
	}

	//validate snubmitted password is the retrievedLogin pw
	validLogin := retrievedLogin.Password == login.Password
	if !validLogin {
		return "", errors.New("invalid login")
	}

	//check if token exists
	token, _ := s.Repo.retrieveToken(login.Email)

	//if no token exists generate
	if token == "" {
		token, err = s.generateToken(login.Email)
		if err != nil {
			return "", fmt.Errorf("%v --> %s", err, "error generating token")
		}
	}

	return token, nil
}

func (s *AuthService) ValidateToken(email, token string) (bool, error) {
	//retrieve token from cache
	retrievedToken, err := s.Repo.retrieveToken(email)
	if err != nil {
		return false, fmt.Errorf("%v --> %s", err, "error retrieving token")
	}

	//compare sentToken and retrievedToken
	if token == retrievedToken {
		return true, nil
	}

	// return false, fmt.Errorf("%s", "could not validate token")
	return false, nil
}

func (s *AuthService) Create(login *Login) (string, error) {
	_, err := s.Repo.createLogin(login)
	if err != nil {
		return "", fmt.Errorf("%v --> %s", err, "error creating login")
	}

	token, err := s.generateToken(login.Email)
	if err != nil {
		return "", fmt.Errorf("%v --> %s", err, "error generating token")
	}

	return token, nil
}

func (s *AuthService) generateToken(email string) (string, error) {
	generatedToken, err := createJWT(email)
	if err != nil {
		return "", fmt.Errorf("%v --> %s", err, "error generating token")
	}
	_, err = s.Repo.storeToken(email, generatedToken)
	if err != nil {
		return "", fmt.Errorf("%v --> %s", err, "error storing token")
	}
	return generatedToken, nil
}
