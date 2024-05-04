package repository

import (
	"database/sql"
	"fmt"
	"messengerApp/internal/app/models"
)

type MessageRepository interface {
	SendMessage(senderID, receiverID int, message string) error
	GetMessages(userID int) ([]*models.Message, error)
	GetMessagesWithPagination(limit int, offset int, sortBy string, sortDir string) ([]*models.Message, error)
	GetMessagesWithFilters(limit int, offset int, sortBy string, sortDir string, filters map[string]interface{}) ([]*models.Message, error)
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) SendMessage(senderID, receiverID int, message string) error {
	query := `INSERT INTO messages (sender_id, receiver_id, message, timestamp) VALUES ($1, $2, $3, NOW())`
	_, err := r.db.Exec(query, senderID, receiverID, message)
	return err
}

func (r *messageRepository) GetMessages(userID int) ([]*models.Message, error) {
	messages := []*models.Message{}
	query := `SELECT id, sender_id, receiver_id, message, timestamp FROM messages WHERE sender_id = $1 OR receiver_id = $1 ORDER BY timestamp DESC`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Message, &m.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}

	return messages, nil
}

func (r *messageRepository) GetMessagesWithPagination(limit int, offset int, sortBy string, sortDir string) ([]*models.Message, error) {
	messages := []*models.Message{}
	query := fmt.Sprintf("SELECT * FROM messages ORDER BY %s %s LIMIT %d OFFSET %d", sortBy, sortDir, limit, offset)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.SenderID, &m.ReceiverID, &m.Message, &m.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}

	return messages, nil
}

func (r *messageRepository) GetMessagesWithFilters(limit, offset int, sortBy, sortDir string, filters map[string]interface{}) ([]*models.Message, error) {
	query := "SELECT * FROM messages WHERE 1=1"
	args := []interface{}{}

	// Add filtering by sender_id
	if senderID, ok := filters["sender_id"].(string); ok && senderID != "" {
		query += " AND sender_id = $1"
		args = append(args, senderID)
	}

	// Adding sorting and pagination
	query += fmt.Sprintf(" ORDER BY %s %s LIMIT $%d OFFSET $%d", sortBy, sortDir, len(args)+1, len(args)+2)
	args = append(args, limit, offset)

	// Querying the database and processing the results
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []*models.Message{}
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Message, &msg.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}

	return messages, nil
}
