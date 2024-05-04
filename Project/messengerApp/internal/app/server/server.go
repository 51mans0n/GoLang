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
}

func NewServer(config *config.ServerConfig, authService service.AuthService, friendService service.FriendService, messageService service.MessageService) *Server {
	router := gin.Default()
	server := &Server{
		router:         router,
		config:         config,
		authService:    authService,
		friendService:  friendService,
		messageService: messageService,
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

	// Routes that do not require authentication
	s.router.POST("/login", authHandler.Login)
	s.router.POST("/register", authHandler.Register)

	authenticated := s.router.Group("/")
	authenticated.Use(authMiddleware)
	{
		adminRoutes := authenticated.Group("/")
		adminRoutes.Use(middleware.AdminAuthMiddleware())

		{
			adminRoutes.GET("/messages", messageHandler.GetMessages)
		}

		authenticated.POST("/add-friend", friendHandler.AddFriend)
		authenticated.POST("/send-message", messageHandler.SendMessage)
	}
}
