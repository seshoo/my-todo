package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
