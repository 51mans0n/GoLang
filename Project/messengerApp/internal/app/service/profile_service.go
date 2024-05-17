package service

import (
	"log"
	"messengerApp/internal/app/models"
	"messengerApp/internal/app/repository"
)

type UpdateProfileRequest struct {
	UserID  int             `json:"user_id"`
	Profile *models.Profile `json:"profile"`
}
type ProfileService interface {
	GetProfile(userID int) (*models.Profile, error)
	UpdateProfile(userID int, profile *models.Profile) error
	CreateProfile(userID int, profile *models.Profile) error
	GetProfilesWithPagination(limit, offset int, sortBy, sortDir string) ([]*models.Profile, error)
	GetProfilesWithFilters(limit, offset int, sortBy, sortDir, filter string) ([]*models.Profile, error)
}

type profileService struct {
	profileRepo repository.ProfileRepository
}

func NewProfileService(profileRepo repository.ProfileRepository) ProfileService {
	return &profileService{
		profileRepo: profileRepo,
	}
}

func (s *profileService) GetProfile(userID int) (*models.Profile, error) {
	profile, err := s.profileRepo.GetProfile(userID)
	if err != nil {
		log.Printf("Service error: %v", err)
		return nil, err
	}

	return profile, nil
}

func (s *profileService) UpdateProfile(userID int, profile *models.Profile) error {
	err := s.profileRepo.UpdateProfile(userID, profile)
	if err != nil {
		return err
	}
	return nil
}

func (s *profileService) CreateProfile(userID int, profile *models.Profile) error {
	err := s.profileRepo.CreateProfile(userID, profile)
	if err != nil {
		return err
	}
	return nil
}

func (s *profileService) GetProfilesWithPagination(limit, offset int, sortBy, sortDir string) ([]*models.Profile, error) {
	return s.profileRepo.GetProfilesWithPagination(limit, offset, sortBy, sortDir)
}

func (s *profileService) GetProfilesWithFilters(limit, offset int, sortBy, sortDir, filter string) ([]*models.Profile, error) {
	return s.profileRepo.GetProfilesWithFilters(limit, offset, sortBy, sortDir, filter)
}
