package domain

type (
	TodoList struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	UsersList struct {
		Id     int
		UserId int
		ListId int
	}

	TodoItem struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Done        bool   `json:"done"`
	}

	ListsItem struct {
		Id     int
		ListId int
		ItemId int
	}
)
