package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"messengerApp/internal/app/models"
)

type ProfileRepository interface {
	UpdateProfile(userID int, profile *models.Profile) error
	GetProfile(userID int) (*models.Profile, error)
	CreateProfile(userID int, profile *models.Profile) error
	GetProfilesWithPagination(limit, offset int, sortBy, sortDir string) ([]*models.Profile, error)
	GetProfilesWithFilters(limit, offset int, sortBy, sortDir, filter string) ([]*models.Profile, error)
}

type profileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) UpdateProfile(userID int, profile *models.Profile) error {
	// Проверка, что соединение с базой данных было инициализировано
	if r.db == nil {
		return errors.New("database connection is not initialized")
	}

	// Проверка, что объект профиля не равен nil
	if profile == nil {
		return errors.New("profile object is nil")
	}

	query := "UPDATE profiles SET name=$1, surname=$2 WHERE user_id=$3"
	_, err := r.db.Exec(query, profile.Name, profile.Surname, userID)
	if err != nil {
		log.Printf("Error updating profile: %v", err)
		return err
	}

	return nil
}

func (r *profileRepository) GetProfile(userID int) (*models.Profile, error) {
	query := "SELECT id, name, surname FROM profiles WHERE user_id=$1"

	row := r.db.QueryRow(query, userID)

	var profile models.Profile

	err := row.Scan(&profile.ID, &profile.Name, &profile.Surname)
	if err != nil {
		log.Printf("Error fetching profile: %v", err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &profile, nil
}

func (r *profileRepository) CreateProfile(userID int, profile *models.Profile) error {
	query := "INSERT INTO profiles (user_id, name, surname) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(query, userID, profile.Name, profile.Surname)
	if err != nil {
		return err
	}

	return nil
}

func (r *profileRepository) GetProfilesWithPagination(limit, offset int, sortBy, sortDir string) ([]*models.Profile, error) {
	query := fmt.Sprintf("SELECT id, user_id, name, surname FROM profiles ORDER BY %s %s LIMIT $1 OFFSET $2", sortBy, sortDir)

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []*models.Profile
	for rows.Next() {
		var profile models.Profile
		if err := rows.Scan(&profile.ID, &profile.UserID, &profile.Name, &profile.Surname); err != nil {
			return nil, err
		}
		profiles = append(profiles, &profile)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r *profileRepository) GetProfilesWithFilters(limit, offset int, sortBy, sortDir, filter string) ([]*models.Profile, error) {
	query := fmt.Sprintf("SELECT id, user_id, name, surname FROM profiles WHERE name ILIKE $3 OR surname ILIKE $3 ORDER BY %s %s LIMIT $1 OFFSET $2", sortBy, sortDir)

	rows, err := r.db.Query(query, limit, offset, "%"+filter+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []*models.Profile
	for rows.Next() {
		var profile models.Profile
		if err := rows.Scan(&profile.ID, &profile.UserID, &profile.Name, &profile.Surname); err != nil {
			return nil, err
		}
		profiles = append(profiles, &profile)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}
