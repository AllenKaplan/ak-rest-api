package auth

import (
	"time"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte("jwt_secret_key")

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CreateJWT func will used to create the JWT while signing in and signing out
func createJWT(userID int, email string) (response string, err error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err == nil {
		return tokenString, nil
	}
	return "", err
}

// VerifyToken func will used to Verify the JWT Token while using APIS
// func verifyToken(tokenString string) (*Claims, error) {
func claimsFromToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if token == nil {
		return nil, err
	}

	return claims, nil

}

var myJwtMiddleware = jwtMiddleware.New(jwtMiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMid := *myJwtMiddleware
		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(401)
		}
		token, _ := jwtMiddleware.FromAuthHeader(c.Request)
		claims, _ := claimsFromToken(token)
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
	}
}
