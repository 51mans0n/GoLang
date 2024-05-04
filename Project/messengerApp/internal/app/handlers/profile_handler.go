package handlers

import (
	"messengerApp/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService service.ProfileService
}

func NewProfileHandler(profileService service.ProfileService) *ProfileHandler {
	return &ProfileHandler{profileService: profileService}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID := c.Param("userID") // Assuming you're passing the userID in the URL path

	// Convert userID to an integer
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call ProfileService's GetProfile method
	profile, err := h.profileService.GetProfile(userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	var updateProfileRequest service.UpdateProfileRequest
	if err := c.BindJSON(&updateProfileRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call ProfileService's UpdateProfile method
	err := h.profileService.UpdateProfile(updateProfileRequest.UserID, updateProfileRequest.Profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
