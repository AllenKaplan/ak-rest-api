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
	jwt, _ := s.Repo.retrieveToken(login.Email)

	//if no token exists generate
	if jwt == "" {
		jwt, err = s.generateToken(retrievedLogin.UserID, retrievedLogin.Email)
		if err != nil {
			return nil, fmt.Errorf("%v --> %s", err, "error generating token")
		}
	}

	token := &Token{
		ID:    retrievedLogin.UserID,
		Email: retrievedLogin.Email,
		Token: jwt,
	}

	return token, nil
}

func (s *AuthService) ValidateToken(email, token string) (*Claims, error) {
	//retrieve token from cache
	retrievedToken, err := s.Repo.retrieveToken(email)
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

	if claims.Email != email {

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

func (s *AuthService) Update(login *Login) (bool, error) {
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
	_, err = s.Repo.storeToken(email, generatedToken)
	if err != nil {
		return "", fmt.Errorf("%v --> %s", err, "error storing token")
	}
	return generatedToken, nil
}
