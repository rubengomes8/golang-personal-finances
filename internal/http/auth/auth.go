package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rubengomes8/golang-personal-finances/internal/repository/models"

	jwt "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenLifespanInHours = 1
	apiSecret            = "unsafeHere" // TODO
)

func EncryptPassword(username, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}

	return string(hash), nil
}

func LoginCheck(ctx context.Context, username, password string, user models.UserTable) (string, error) {

	err := verifyPassword(password, user.Passhash)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", fmt.Errorf("invalid password: %v", err)
	}

	token, err := generateToken(uint(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil

}

func verifyPassword(password, hashedPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password))
}

func generateToken(userID uint) (string, error) {

	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespanInHours)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(apiSecret))

}

func validateToken(ctx *gin.Context) error {

	tokenString := extractToken(ctx)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apiSecret), nil
	})

	return err
}

func extractToken(ctx *gin.Context) string {

	token := ctx.Query("token")
	if token != "" {
		return token
	}

	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}
