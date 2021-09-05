package domain

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"usermane"`
	Password string `json:"password"`
}
