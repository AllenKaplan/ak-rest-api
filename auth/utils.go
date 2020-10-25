package auth

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func generateToken(userID int) *Token {
	return &Token{
		TokenID: rand.Intn(100000) + 1000000,
		UserID:  userID,
		Expiry:  0,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
