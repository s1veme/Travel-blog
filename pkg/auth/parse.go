package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
)

func ParseToken(accessToken string, signingKey []byte) (string, error) {
	fmt.Println(accessToken)
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Email, nil
	}

	return "", ErrInvalidAccessToken
}
