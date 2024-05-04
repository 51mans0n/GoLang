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

func (s *friendService) GetFriends(userID int) ([]*models.User, error) {
	var friends []*models.User

	rows, err := s.db.Query(`
        SELECT u.id, u.username 
        FROM users u
        JOIN friends f ON f.friend_id = u.id
        WHERE f.user_id = $1`, userID)
	if err != nil {
		log.Printf("Error retrieving friends: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			log.Printf("Error scanning friend: %v", err)
			continue
		}
		friends = append(friends, &user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		return nil, err
	}

	return friends, nil
}

func (s *friendService) SendMessage(senderID, receiverID int, message string) error {
	// Simulate sending message by returning nil
	return nil
}
