package auth

import (
	"blog/internal/applications/model"
)

type UseCase interface {
	SignIn(user *model.User) (string, error)
}
