package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	Db *sql.DB
}

func New(db *sql.DB) Store {
	return Store{
		Db: db,
	}
}
