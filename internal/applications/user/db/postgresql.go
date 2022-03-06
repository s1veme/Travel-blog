package user

import (
	"blog/internal/applications/store"
	"blog/internal/applications/user"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type repository struct {
	store  store.Store
	logger *logrus.Logger
}

func NewRepository(client store.Store, logger *logrus.Logger) user.UserRepository {

	return &repository{
		store:  client,
		logger: logger,
	}
}

func (r *repository) Create(u *user.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.Db.QueryRow(
		"INSERT INTO users (email, username, encrypted_password) VALUES ($1, $2, $3) RETURNING id",
		u.Email, u.Username, u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *repository) FindByEmail(email string) (*user.User, error) {
	u := &user.User{}
	if err := r.store.Db.QueryRow(
		"SELECT id, email, username, encrypted_password FROM users WHERE email = $1",
		email).Scan(&u.ID, &u.Email, &u.Username, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}
