package db

import (
	"blog/internal/applications/model"
	"blog/internal/applications/post"
	"blog/internal/applications/store"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type repository struct {
	store  store.Store
	logger *logrus.Logger
}

func NewRepository(client store.Store, logger *logrus.Logger) post.PostRepository {
	return &repository{
		store:  client,
		logger: logger,
	}
}

func (r *repository) Create(m *model.Post) error {
	return r.store.Db.QueryRow(
		"INSERT INTO posts (title, content, owner) VALUES ($1, $2, $3) RETURNING id",
		m.Title, m.Content, m.Owner,
	).Scan(&m.ID)
}

func (r *repository) GetList() (*[]model.Post, error) {
	rows, err := r.store.Db.Query("SELECT id, title, content, owner FROM posts")

	posts := make([]model.Post, 0)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p model.Post

		err = rows.Scan(&p.ID, &p.Title, &p.Content, &p.Owner)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (r *repository) GetByID(id int) (*model.Post, error) {
	p := &model.Post{}

	err := r.store.Db.QueryRow(
		"SELECT id, title, content, owner FROM posts WHERE id = $1", id).Scan(&p.ID, &p.Title, &p.Content, &p.Owner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	_, err := r.store.Db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
