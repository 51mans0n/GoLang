package repository

import (
	"database/sql"
	"messengerApp/internal/app/models"
)

type ProfileRepository interface {
	UpdateProfile(userID int, profile *models.Profile) error
	GetProfile(userID int) (*models.Profile, error)
}

type profileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) UpdateProfile(userID int, profile *models.Profile) error {
	// Implementation for updating a user's profile
	return nil
}

func (r *profileRepository) GetProfile(userID int) (*models.Profile, error) {
	// Implementation for getting a user's profile
	return nil, nil
}
