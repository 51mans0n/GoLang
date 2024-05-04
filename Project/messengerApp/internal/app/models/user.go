package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []string
}

func (u *User) IsAdmin() bool {
	for _, role := range u.Roles {
		if role == "admin" {
			return true
		}
	}
	return false
}
