package service

import (
	"database/sql"
	"errors"
	"log"
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
	"messengerApp/internal/utils"
)

type AuthService interface {
	Login(username, password string) (string, error)
	Register(username, password string) error
	RegisterUser(username, password string) (*models.User, error)
	DeleteUser(id string) error
	UpdateUser(id string, user models.User) error
	UserRepo() repository.UserRepository
}

type authService struct {
	userRepo repository.UserRepository
	db       *sql.DB
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *authService) UserRepo() repository.UserRepository {
	return s.userRepo
}

// NewAuthService creates a new AuthService instance
func NewAuthService(userRepo repository.UserRepository, db *sql.DB) AuthService {
	return &authService{userRepo: userRepo, db: db}
}

// Login handles user login
func (s *authService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if user == nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}

	// Generating a token for the user
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Register registers a new user
func (s *authService) Register(username, password string) error {
	user := &models.User{Username: username, Password: password}
	return s.userRepo.Create(user)
}

// RegisterUser registers a user and returns a user object
func (s *authService) RegisterUser(username, password string) (*models.User, error) {
	user := &models.User{Username: username, Password: password}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// loadUserRoles loads user roles from the database
func (s *authService) loadUserRoles(userID int) []string {
	var roles []string
	rows, err := s.db.Query("SELECT r.name FROM roles r JOIN user_roles ur ON r.id = ur.role_id WHERE ur.user_id = $1", userID)
	if err != nil {
		log.Printf("Failed to load user roles: %v\n", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			log.Printf("Failed to scan role: %v\n", err)
			continue
		}
		roles = append(roles, role)
	}

	return roles
}

// DeleteUser deletes a user by ID.
func (s *authService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}

// UpdateUser updates user data by ID.
func (s *authService) UpdateUser(id string, user models.User) error {
	return s.userRepo.Update(id, user)
}
