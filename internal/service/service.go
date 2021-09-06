package service

import "github.com/seshoo/my-todo/internal/repository"

type (
	User interface {
	}

	TodoList interface {
	}

	TodoItem interface {
	}

	Service struct {
		User
		TodoList
		TodoItem
	}

	Dependencies struct {
		Repository *repository.Repository
	}
)

func NewService(d Dependencies) *Service {
	return &Service{}
}
