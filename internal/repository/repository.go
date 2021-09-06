package repository

import "github.com/jmoiron/sqlx"

type (
	User interface {
	}

	TodoList interface {
	}

	TodoItem interface {
	}

	Repository struct {
		User
		TodoList
		TodoItem
	}
)

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
