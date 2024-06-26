package server

import (
	"messengerApp/cmd/config"
	"messengerApp/internal/app/handlers"
	"messengerApp/internal/app/middleware"
	"messengerApp/internal/app/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router         *gin.Engine
	config         *config.ServerConfig
	authService    service.AuthService
	friendService  service.FriendService
	messageService service.MessageService
	profileHandler *handlers.ProfileHandler
}

func NewServer(config *config.ServerConfig, authService service.AuthService, friendService service.FriendService, messageService service.MessageService, profileHandler *handlers.ProfileHandler) *Server {
	router := gin.Default()
	server := &Server{
		router:         router,
		config:         config,
		authService:    authService,
		friendService:  friendService,
		messageService: messageService,
		profileHandler: profileHandler,
	}
	server.setupRoutes()
	return server
}

func (s *Server) Run() {
	if s.config == nil || s.config.Port == "" {
		panic("Server configuration is missing or invalid")
	}
	s.router.Run(":" + s.config.Port)
}

func (s *Server) setupRoutes() {
	authHandler := handlers.NewAuthHandler(s.authService)
	friendHandler := handlers.NewFriendHandler(s.friendService)
	messageHandler := handlers.NewMessageHandler(s.messageService)
	authMiddleware := middleware.NewAuthMiddleware(s.authService.UserRepo())
	adminMiddleware := middleware.AdminAuthMiddleware()

	// Routes that do not require authentication
	s.router.POST("/login", authHandler.Login)
	s.router.POST("/register", authHandler.Register)

	authenticated := s.router.Group("/")
	authenticated.Use(authMiddleware)
	{
		adminRoutes := authenticated.Group("/")
		adminRoutes.Use(adminMiddleware)
		{
			adminRoutes.GET("/messages", messageHandler.GetMessages)
			adminRoutes.DELETE("/users/:id", authHandler.DeleteUser)
			adminRoutes.PUT("/users/:id", authHandler.UpdateUser)
			adminRoutes.GET("/friends", friendHandler.GetFriends)
			adminRoutes.GET("/profiles/:userID", s.profileHandler.GetProfile)
			adminRoutes.PUT("/profiles/update-profile/:userID", s.profileHandler.UpdateProfile)
			adminRoutes.POST("/profiles/create-profile/:userID", s.profileHandler.CreateProfile)
			adminRoutes.GET("/profiles", s.profileHandler.GetProfiles)
		}
		authenticated.POST("/add-friend", friendHandler.AddFriend)
		authenticated.POST("/send-message", messageHandler.SendMessage)
	}
}
