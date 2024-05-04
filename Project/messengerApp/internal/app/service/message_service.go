package service

import (
	"database/sql"
	"errors"
	"log"
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
)

type SendMessageRequest struct {
	UserID   int    `json:"user_id"`
	FriendID int    `json:"friend_id"`
	Message  string `json:"message"`
}
type MessageService interface {
	SendMessage(senderID, receiverID int, message string) error
	GetMessages(userID int) ([]*models.Message, error)
	GetMessagesWithPagination(page int, pageSize int, sortBy string, sortDir string) ([]*models.Message, error)
	GetMessagesWithFilters(page, pageSize int, sortBy, sortDir, senderID string) ([]*models.Message, error)
}

type messageService struct {
	db          *sql.DB
	userRepo    repository.UserRepository
	friendRepo  repository.FriendRepository
	messageRepo repository.MessageRepository
}

func NewMessageService(db *sql.DB, userRepo repository.UserRepository, friendRepo repository.FriendRepository, messageRepo repository.MessageRepository) MessageService {
	return &messageService{
		db:          db,
		userRepo:    userRepo,
		friendRepo:  friendRepo,
		messageRepo: messageRepo,
	}
}

func (s *messageService) SendMessage(senderID, receiverID int, message string) error {
	if senderID == 0 || receiverID == 0 {
		log.Println("Invalid sender or receiver ID")
		return errors.New("invalid sender or receiver ID")
	}

	_, err := s.db.Exec("INSERT INTO messages (sender_id, receiver_id, message) VALUES ($1, $2, $3)", senderID, receiverID, message)
	if err != nil {
		log.Printf("Error sending message: %v", err)
		return err
	}

	log.Println("Message sent successfully")
	return nil
}

func (s *messageService) GetMessages(userID int) ([]*models.Message, error) {
	return s.messageRepo.GetMessages(userID)
}

func (s *messageService) GetMessagesWithPagination(page int, pageSize int, sortBy string, sortDir string) ([]*models.Message, error) {
	offset := (page - 1) * pageSize
	return s.messageRepo.GetMessagesWithPagination(pageSize, offset, sortBy, sortDir)
}

func (s *messageService) GetMessagesWithFilters(page, pageSize int, sortBy, sortDir, senderID string) ([]*models.Message, error) {
	filter := map[string]interface{}{}
	if senderID != "" {
		filter["sender_id"] = senderID
	}

	// Calling the repository with passing all parameters and handling a possible error
	messages, err := s.messageRepo.GetMessagesWithFilters(pageSize, (page-1)*pageSize, sortBy, sortDir, filter)
	if err != nil {
		return nil, err // Убедитесь, что возвращаете ошибку
	}

	return messages, nil
}
