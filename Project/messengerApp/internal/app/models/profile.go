package models

type Profile struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
