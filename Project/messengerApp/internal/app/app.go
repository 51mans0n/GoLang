package app

import (
	"errors"

	"database/sql"
	"messengerApp/cmd/config" // Adjust import path if necessary
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
	"messengerApp/internal/app/service"
	"messengerApp/internal/utils"
)

type App struct {
	db             *sql.DB
	userRepo       repository.UserRepository
	friendRepo     repository.FriendRepository
	messageRepo    repository.MessageRepository
	profileRepo    repository.ProfileRepository
	authService    service.AuthService
	friendService  service.FriendService
	messageService service.MessageService
	profileService service.ProfileService
}

// Login implements service.AuthService.
func (a *App) Login(username string, password string) (string, error) {
	// Retrieve the user by username from the repository
	user, err := a.userRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	// Check if the user exists and if the password matches
	if user == nil || user.Password != password {
		return "", errors.New("invalid credentials")
	}

	// Generate a token for the user
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Register implements service.AuthService.
func (a *App) Register(username string, password string) error {
	// Check if the username already exists in the repository
	existingUser, err := a.userRepo.FindByUsername(username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Create a new user with the provided username and password
	user := &models.User{Username: username, Password: password}
	err = a.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

// RegisterUser implements service.AuthService.
func (a *App) RegisterUser(username string, password string) (*models.User, error) {
	// Check if the username already exists in the repository
	existingUser, err := a.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// Create a new user with the provided username and password
	user := &models.User{Username: username, Password: password}
	err = a.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewApp(db *sql.DB, cfg *config.ServerConfig) *App {
	// Create user repository
	userRepo := repository.NewUserRepository(db)

	// Create other repositories
	friendRepo := repository.NewFriendRepository(db)
	messageRepo := repository.NewMessageRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	// Create service instances
	authService := service.NewAuthService(userRepo, db)
	friendService := service.NewFriendService(db, userRepo, friendRepo)
	messageService := service.NewMessageService(db, userRepo, friendRepo, messageRepo)
	profileService := service.NewProfileService(userRepo, profileRepo)

	return &App{
		db:             db,
		userRepo:       userRepo,
		friendRepo:     friendRepo,
		messageRepo:    messageRepo,
		profileRepo:    profileRepo,
		authService:    authService,
		friendService:  friendService,
		messageService: messageService,
		profileService: profileService,
	}
}

func (a *App) AuthService() service.AuthService {
	return a.authService
}

func (a *App) FriendService() service.FriendService {
	return a.friendService
}

func (a *App) MessageService() service.MessageService {
	return a.messageService
}

func (a *App) ProfileService() service.ProfileService {
	return a.profileService
}
