package service

import (
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
)

type ProfileService interface {
	GetProfile(userID int) (*models.Profile, error)
	UpdateProfile(userID int, profile *models.Profile) error
}
type UpdateProfileRequest struct {
	UserID  int             `json:"user_id"`
	Profile *models.Profile `json:"profile"`
}
type profileService struct {
	userRepo    repository.UserRepository
	profileRepo repository.ProfileRepository
}

func NewProfileService(userRepo repository.UserRepository, profileRepo repository.ProfileRepository) ProfileService {
	return &profileService{
		userRepo:    userRepo,
		profileRepo: profileRepo,
	}
}

func (s *profileService) GetProfile(userID int) (*models.Profile, error) {
	// Simulate getting profile by returning empty profile and no error
	return &models.Profile{}, nil
}

func (s *profileService) UpdateProfile(userID int, profile *models.Profile) error {
	// Simulate updating profile by returning nil
	return nil
}
