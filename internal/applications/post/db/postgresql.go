package db

import (
	"blog/internal/applications/model"
	"blog/internal/applications/post"
	"blog/internal/applications/store"
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
	panic("implement me")
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

func (r *repository) GetByID(id int) (model.Post, error) {
	panic("implement me")
}

func (r *repository) Delete(id int) error {
	panic("implement me")
}
