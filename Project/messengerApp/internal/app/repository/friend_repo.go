package repository

import (
	"database/sql"
	"messengerApp/internal/app/models"
)

type FriendRepository interface {
	AddFriend(userID, friendID int) error
	GetFriends(userID int) ([]*models.User, error)
}

type friendRepository struct {
	db *sql.DB
}

func NewFriendRepository(db *sql.DB) FriendRepository {
	return &friendRepository{db: db}
}

func (r *friendRepository) AddFriend(userID, friendID int) error {
	// Implementation for adding a friend
	return nil
}

func (r *friendRepository) GetFriends(userID int) ([]*models.User, error) {
	// Implementation for getting friends of a user
	return nil, nil
}
