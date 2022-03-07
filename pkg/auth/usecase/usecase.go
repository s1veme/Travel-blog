package usecase

import (
	"blog/internal/applications/model"
	"blog/pkg/auth"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type Authorizer struct {
	repository     auth.Repository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthorizer(repository auth.Repository, hashSalt string, signingKey []byte, expireDuration time.Duration) *Authorizer {
	return &Authorizer{
		repository:     repository,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

func (a *Authorizer) SignIn(user *model.User) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(a.hashSalt))
	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Email: user.Email,
	})

	return token.SignedString(a.signingKey)
}
