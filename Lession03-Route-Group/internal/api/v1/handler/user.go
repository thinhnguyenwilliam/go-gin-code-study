package v1handler

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/utils"
)

type UserHandler struct {
	validate *validator.Validate
}

func NewUserHandler() *UserHandler {
	v := validator.New()
	_ = v.RegisterValidation("alphanumspace", utils.AlphaNumSpace)
	return &UserHandler{validate: v}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of all users (V1)",
		"users":   []string{"Alice", "Bob", "Charlie"},
	})
}

func (h *UserHandler) GetUserWithoutSlug(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"slug": "no news",
	})
}

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

func (h *UserHandler) GetUserBySlug(c *gin.Context) {
	slug := c.Param("slug")

	// Validate slug format
	if !slugRegex.MatchString(slug) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid slug",
			"message": "Slug must contain only lowercase letters, numbers, and hyphens",
		})
		return
	}

	// Continue with lookup logic
	c.JSON(http.StatusOK, gin.H{
		"type":    "Slug User",
		"message": "User details for slug: " + slug,
	})
}

func (h *UserHandler) GetUsersByUUID(c *gin.Context) {
	id := c.Param("uuid")

	// Validate UUID format
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid UUID",
			"message": "UUID must be in the correct format",
		})
		return
	}

	// Continue with logic (e.g., lookup from DB)
	c.JSON(http.StatusOK, gin.H{
		"ID":      id,
		"type":    "UUID User",
		"message": "User details for UUID " + id,
	})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	// Validate id: check if it's a number and positive
	userID, err := strconv.Atoi(id)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"User ID": userID,
		"message": "User details for ID " + strconv.Itoa(userID),
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "New user created",
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated user with ID " + id,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted user with ID " + id,
	})
}
