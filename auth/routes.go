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
	jwt, err := s.Repo.retrieveToken(retrievedLogin.UserID)

	//if no token exists generate
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("%v --> %s", err, "auth.routes.Login")
		}
	}

	token := &Token{
		ID:    retrievedLogin.UserID,
		Email: retrievedLogin.Email,
		Token: jwt,
	}

	return token, nil
}

func (s *AuthService) ValidateToken(userID int, token string) (*Claims, error) {
	//retrieve token from cache
	retrievedToken, err := s.Repo.retrieveToken(userID)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error retrieving token")
	}

	//compare sentToken and retrievedToken
	if token != retrievedToken {
		return nil, errors.New("Token not the same as previously stored; may be expired")
	}

	//get claims
	claims, err := claimsFromToken(token)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error retrieving claims")
	}

	if userID != claims.UserID {
		return nil, errors.New("UserID different than token claim")
	}

	return claims, nil
}

func (s *AuthService) Create(login *Login) (*Token, error) {
	retrievedLogin, err := s.Repo.create(login)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error creating login")
	}

	jwt, err := s.generateToken(login.UserID, login.Email)
	if err != nil {
		return nil, fmt.Errorf("%v --> %s", err, "error generating token")
	}

	token := &Token{
		ID:    retrievedLogin.UserID,
		Email: retrievedLogin.Email,
		Token: jwt,
	}

	return token, nil
}

func (s *AuthService) Update(userID int, loginRequest *LoginRequest) (bool, error) {
	login := &Login{
		UserID:   userID,
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}

	updateSuccess, err := s.Repo.update(login)
	if err != nil {
		return false, fmt.Errorf("%v --> %s", err, "error creating login")
	}

	return updateSuccess, nil
}

func (s *AuthService) generateToken(userID int, email string) (string, error) {
	generatedToken, err := createJWT(userID, email)
	if err != nil {
		return "", fmt.Errorf("%v --> %s", err, "error generating token")
	}
	_, err = s.Repo.storeToken(userID, generatedToken)
	if err != nil {
		return "", fmt.Errorf("%v --> %s", err, "error storing token")
	}
	return generatedToken, nil
}
