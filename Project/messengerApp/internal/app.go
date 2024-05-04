package app

import (
	"database/sql"
	"messengerApp/cmd/config"
	"messengerApp/internal/app/repository"
	"messengerApp/internal/app/service"
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

func NewApp(db *sql.DB, cfg *config.ServerConfig) *App {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	friendRepo := repository.NewFriendRepository(db)
	messageRepo := repository.NewMessageRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, db)
	friendService := service.NewFriendService(db, userRepo, friendRepo)
	messageService := service.NewMessageService(db, userRepo, friendRepo, messageRepo)
	profileService := service.NewProfileService(userRepo, profileRepo)

	// Return the initialized app
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
