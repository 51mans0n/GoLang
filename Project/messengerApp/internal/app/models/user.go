package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []Role `json:"roles"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u *User) IsAdmin() bool {
	for _, role := range u.Roles {
		if role.Name == "admin" {
			return true
		}
	}
	return false
}
