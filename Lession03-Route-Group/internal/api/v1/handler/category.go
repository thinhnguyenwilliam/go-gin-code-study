package v1handler

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/dto"
	"github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/utils"
)

type CategoryHandler struct {
	validate *validator.Validate
}

func NewCategoryHandler() *CategoryHandler {
	v := validator.New()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("alphanumspace", utils.AlphaNumSpace)
		_ = v.RegisterValidation("imgext", utils.ValidateImageExtension)
	}
	return &CategoryHandler{validate: v}
}

func (h *CategoryHandler) UploadCategoryImage(c *gin.Context) {
	// Bind query parameters
	var query dto.UploadCategoryQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid query parameters",
			"fields": utils.FormatValidationErrors(err),
		})
		return
	}

	// Bind form fields
	var form dto.UploadCategoryForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Validation failed",
			"fields": utils.FormatValidationErrors(err),
		})
		return
	}

	// File validation
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	if !allowed[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unsupported file type. Allowed: jpg, jpeg, png",
		})
		return
	}

	uploadPath := "uploads/categories"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	dst := filepath.Join(uploadPath, file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "File uploaded successfully",
		"name":        form.Name,
		"description": form.Description,
		"user_id":     query.UserID,
		"source":      query.Source,
		"file":        file.Filename,
		"path":        dst,
	})
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest

	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		c.JSON(400, gin.H{
			"error":  "Validation failed",
			"fields": utils.FormatValidationErrors(err),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Category created successfully",
		"data":    req,
	})
}
