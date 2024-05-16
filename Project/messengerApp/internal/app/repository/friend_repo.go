package repository

import (
	"database/sql"
	"fmt"
	"log"
	"messengerApp/internal/app/models"
)

type FriendRepository interface {
	AddFriend(userID, friendID int) error
	GetFriends(userID int) ([]*models.User, error)
	GetFriendsWithPagination(limit, offset int, sortBy, sortDir string) ([]*models.User, error)
	GetFriendsWithFilters(limit, offset int, sortBy, sortDir, filter string) ([]*models.User, error)
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

func (r *friendRepository) GetFriendsWithPagination(limit, offset int, sortBy, sortDir string) ([]*models.User, error) {
	query := fmt.Sprintf("SELECT id, username FROM users ORDER BY %s %s LIMIT $1 OFFSET $2", sortBy, sortDir)
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}
		friends = append(friends, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return friends, nil
}

// GetFriendsWithFilters retrieves a filtered and paginated list of friends for a user.
func (r *friendRepository) GetFriendsWithFilters(limit, offset int, sortBy, sortDir, filter string) ([]*models.User, error) {
	query := fmt.Sprintf("SELECT id, username FROM users WHERE username LIKE '%%%s%%' ORDER BY %s %s LIMIT $1 OFFSET $2", filter, sortBy, sortDir)
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}
		friends = append(friends, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return friends, nil
}

func (r *friendRepository) GetFriends(userID int) ([]*models.User, error) {
	var friends []*models.User

	rows, err := r.db.Query(`
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
