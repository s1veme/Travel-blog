package post

import "blog/internal/applications/model"

type PostRepository interface {
	Create(*model.Post) error
	GetList() (*[]model.Post, error)
	// GetByID(id int) (model.Post, error)
	// Delete(id int) error
}
