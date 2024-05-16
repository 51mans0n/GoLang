package service

import (
	"database/sql"
	"errors"
	"log"
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
)

type FriendService interface {
	AddFriend(userID, friendID int) error
	GetFriends(userID int) ([]*models.User, error)
	SendMessage(senderID, receiverID int, message string) error
	GetFriendsWithPagination(page, pageSize int, sortBy, sortDir string) ([]*models.User, error)
	GetFriendsWithFilters(page, pageSize int, sortBy, sortDir, filter string) ([]*models.User, error)
}

type friendService struct {
	db         *sql.DB
	userRepo   repository.UserRepository
	friendRepo repository.FriendRepository
}

func NewFriendService(db *sql.DB, userRepo repository.UserRepository, friendRepo repository.FriendRepository) FriendService {
	return &friendService{
		db:         db,
		userRepo:   userRepo,
		friendRepo: friendRepo,
	}
}

func (s *friendService) AddFriend(userID, friendID int) error {
	if userID == friendID {
		log.Println("Attempt to add self as friend")
		return errors.New("cannot add yourself as a friend")
	}

	exists, err := s.friendExists(userID, friendID)
	if err != nil {
		log.Printf("Error checking if friendship exists: %v", err)
		return err
	}
	if exists {
		log.Println("Users are already friends")
		return errors.New("these users are already friends")
	}

	_, err = s.db.Exec("INSERT INTO friends (user_id, friend_id) VALUES ($1, $2)", userID, friendID)
	if err != nil {
		log.Printf("Error adding friend: %v", err)
		return err
	}
	return nil
}

func (s *friendService) friendExists(userID, friendID int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(
                 SELECT 1 
                 FROM friends 
                 WHERE (user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1)
              )`
	err := s.db.QueryRow(query, userID, friendID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *friendService) SendMessage(senderID, receiverID int, message string) error {
	// Simulate sending message by returning nil
	return nil
}
func (s *friendService) GetFriendsWithPagination(page, pageSize int, sortBy, sortDir string) ([]*models.User, error) {
	offset := (page - 1) * pageSize
	return s.friendRepo.GetFriendsWithPagination(pageSize, offset, sortBy, sortDir)
}

func (s *friendService) GetFriendsWithFilters(page, pageSize int, sortBy, sortDir, filter string) ([]*models.User, error) {
	return s.friendRepo.GetFriendsWithFilters(page, pageSize, sortBy, sortDir, filter)
}
func (s *friendService) GetFriends(userID int) ([]*models.User, error) {
	return s.friendRepo.GetFriends(userID)
}
