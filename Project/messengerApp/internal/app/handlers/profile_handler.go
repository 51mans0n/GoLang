package handlers

import (
	"log"
	"net/http"
	"strconv"

	"messengerApp/internal/app/models"
	"messengerApp/internal/app/service"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService service.ProfileService
}

func NewProfileHandler(profileService service.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID := c.Param("userID")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	profile, err := h.profileService.GetProfile(userIDInt)
	if err != nil {
		log.Printf("Fetching profile for userID: %d", userIDInt)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	var req struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Convert userID from string to int
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.profileService.UpdateProfile(userID, &models.Profile{Name: req.Name, Surname: req.Surname})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	userID := c.Param("userID")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var createProfileRequest models.Profile
	if err := c.BindJSON(&createProfileRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err = h.profileService.CreateProfile(userIDInt, &createProfileRequest)
	if err != nil {
		// Log the specific error
		log.Printf("Failed to create profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Profile created successfully"})
}

func (h *ProfileHandler) GetProfiles(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	sortBy := c.Query("sortBy")
	sortDir := c.Query("sortDir")
	filter := c.Query("filter")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10 // default limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0 // default offset
	}

	if sortBy == "" {
		sortBy = "name" // default sort by
	}

	if sortDir == "" {
		sortDir = "asc" // default sort direction
	}

	var profiles []*models.Profile

	if filter != "" {
		profiles, err = h.profileService.GetProfilesWithFilters(limit, offset, sortBy, sortDir, filter)
	} else {
		profiles, err = h.profileService.GetProfilesWithPagination(limit, offset, sortBy, sortDir)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profiles"})
		return
	}

	c.JSON(http.StatusOK, profiles)
}
