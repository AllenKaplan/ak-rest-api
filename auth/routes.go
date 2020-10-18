package auth

func (s AuthService) Login(login *Login) (*Token, error) {

	token := &Token{
		tokenID: 0,
		userID:  0,
		expiry:  0,
	}

	return token, nil
}

func (s AuthService) Validate(token *Token) (error, *Token) {
	return nil, token
}
