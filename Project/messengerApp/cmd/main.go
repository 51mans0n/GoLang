package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"messengerApp/cmd/config"
	"messengerApp/internal/app/repository"
	"messengerApp/internal/app/server"
	"messengerApp/internal/app/service"
	database "messengerApp/internal/database/postgre"

	_ "github.com/lib/pq"
)

func main() {
	// Load application configuration
	dbConfig, serverConfig, err := config.LoadConfig()
	gin.SetMode(gin.ReleaseMode)
	if err != nil {
		fmt.Printf("error loading config: %v\n", err)
		return
	}

	// Initialize the database connection
	db, err := database.InitDatabase(dbConfig)
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
		return
	}
	if db == nil {
		log.Fatal("Database object is nil")
	}
	sqlDB := db.DB()

	// Pass the db.DB() to NewUserRepository
	userRepo := repository.NewUserRepository(sqlDB)
	friendRepo := repository.NewFriendRepository(sqlDB)

	// Initialize the authentication service
	authService := service.NewAuthService(userRepo, sqlDB)
	friendService := service.NewFriendService(sqlDB, userRepo, friendRepo)

	messageRepo := repository.NewMessageRepository(sqlDB)
	messageService := service.NewMessageService(sqlDB, userRepo, friendRepo, messageRepo)

	// Start the HTTP server
	srv := server.NewServer(serverConfig, authService, friendService, messageService)
	srv.Run()
}
