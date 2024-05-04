package repository

import (
	"database/sql"
	"messengerApp/internal/app/models"
)

type UserRepository interface {
	FindByID(id int) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(id int) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE id = $1"
	var user models.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = $1"
	var user models.User
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) AddRole(userID int, role string) error {
	roleIDQuery := "SELECT id FROM roles WHERE name = $1"
	var roleID int
	err := r.db.QueryRow(roleIDQuery, role).Scan(&roleID)
	if err != nil {
		return err
	}

	query := "INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)"
	_, err = r.db.Exec(query, userID, roleID)
	if err != nil {
		return err
	}
	return nil
}
