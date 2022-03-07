package auth

import (
	"blog/internal/applications/model"
)

type Repository interface {
	FindByEmail(string) (*model.User, error)
}
