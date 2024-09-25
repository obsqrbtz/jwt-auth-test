package tokens

import (
	"fmt"
	"jwt-auth-test/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserClaims struct {
	jwt.StandardClaims
	ClientIP string `json:"client_ip"`
	ClientID string `json:"client_id"`
}

func CreateAccessToken(claims *UserClaims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("token.SignedString: %w", err)
	}
	return signedToken, nil
}

func CreateRefreshToken(user *models.User) (string, error) {
	refreshToken, err := bcrypt.GenerateFromPassword([]byte(user.ID), 14)

	if err != nil {
		return "", fmt.Errorf("bcrypt: %w", err)
	}

	return string(refreshToken), nil
}

func CreateUserClaims(user models.User) UserClaims {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.NewString(),
			Issuer:    user.ID,
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
		ClientIP: user.IP,
		ClientID: user.ID,
	}
	return claims
}

func ParseToken(rawToken string, secretKet string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(rawToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKet), nil
	})

	return token.Claims.(*UserClaims), err
}
