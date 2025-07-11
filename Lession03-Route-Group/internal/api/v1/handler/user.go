package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/dto"
	"github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/utils"
)

type UserHandler struct {
	validate *validator.Validate
}

func NewUserHandler() *UserHandler {
	v := validator.New()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("alphanumspace", utils.AlphaNumSpace)
		_ = v.RegisterValidation("slug", utils.ValidateSlug)
	}

	return &UserHandler{validate: v}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	var uri dto.UserQuery

	if err := c.ShouldBindUri(&uri); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "Validation failed",
				"fields": utils.FormatValidationErrors(ve),
			})
			return
		}
		// Handle parse error (e.g. string instead of int)
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID must be a valid positive integer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      uri.ID,
		"message": "User ID is valid",
	})
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

func (h *UserHandler) GetUserBySlug(c *gin.Context) {
	var uri dto.UserSlugQuery

	if err := c.ShouldBindUri(&uri); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "Validation failed",
				"fields": utils.FormatValidationErrors(ve),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"type":    "Slug User",
		"message": "User details for slug: " + uri.Slug,
	})
}

func (h *UserHandler) GetUserByUUID(c *gin.Context) {
	var uri dto.UserUUIDQuery

	if err := c.ShouldBindUri(&uri); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "Validation failed",
				"fields": utils.FormatValidationErrors(ve),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid path parameter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uuid":    uri.UUID,
		"message": "Valid UUID",
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
