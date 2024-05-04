package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
