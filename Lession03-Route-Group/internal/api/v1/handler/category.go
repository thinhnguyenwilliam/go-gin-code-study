package v1handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

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

func (h *CategoryHandler) UploadMultipleCategoryImages(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form"})
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	const maxImages = 5
	if len(files) > maxImages {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Maximum %d images are allowed", maxImages),
		})
		return
	}

	uploadPath := "uploads/categories"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	var uploadedFiles []string
	var failedFiles []map[string]string

	for _, fileHeader := range files {
		fileName, err := validateAndSaveImage(fileHeader, uploadPath)
		if err != nil {
			failedFiles = append(failedFiles, map[string]string{
				"file":  fileHeader.Filename,
				"error": err.Error(),
			})
			continue
		}
		uploadedFiles = append(uploadedFiles, fileName)
	}

	if len(uploadedFiles) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "No valid images uploaded",
			"failed": failedFiles,
		})
		return
	}

	var uploadedURLs []string
	for _, name := range uploadedFiles {
		uploadedURLs = append(uploadedURLs, fmt.Sprintf("/static/categories/%s", name))
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Some or all files uploaded successfully",
		//"files":   uploadedFiles,
		"files":  uploadedURLs,
		"failed": failedFiles,
	})
}

func validateAndSaveImage(fileHeader *multipart.FileHeader, uploadPath string) (string, error) {
	const maxSize = 2 << 20 // 2MB

	if fileHeader.Size > maxSize {
		return "", fmt.Errorf("file too large")
	}

	if err := utils.ValidateFileExtension(fileHeader.Filename, utils.AllowedExtensions); err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	if err := utils.ValidateImageMIME(file); err != nil {
		return "", err
	}

	// Reset the reader
	_, _ = file.Seek(0, 0)

	newName := utils.GenerateUUIDFileName(fileHeader.Filename)
	dst := filepath.Join(uploadPath, newName)

	if err := saveFile(fileHeader, dst); err != nil {
		return "", err
	}

	return newName, nil
}

func saveFile(fileHeader *multipart.FileHeader, dst string) error {
	return os.WriteFile(dst, getFileBytes(fileHeader), os.ModePerm)
}

// Optional helper to read bytes (or reuse gin's SaveUploadedFile if preferred)
func getFileBytes(fileHeader *multipart.FileHeader) []byte {
	file, _ := fileHeader.Open()
	defer file.Close()
	content, _ := io.ReadAll(file)
	return content
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
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open uploaded file"})
		return
	}
	defer file.Close()

	// ✅ Check file extension
	if err := utils.ValidateFileExtension(fileHeader.Filename, utils.AllowedExtensions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ Check MIME content
	if err := utils.ValidateImageMIME(file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ✅ Check file size
	const maxSize = 2 << 20 // 2MB
	if err := utils.ValidateFileSize(fileHeader.Size, maxSize); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadPath := "uploads/categories"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	// ✅ Generate UUID file name
	newFileName := utils.GenerateUUIDFileName(fileHeader.Filename)
	dst := filepath.Join(uploadPath, newFileName)

	if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "File uploaded successfully",
		"name":        form.Name,
		"description": form.Description,
		"user_id":     query.UserID,
		"source":      query.Source,
		"file":        newFileName,
		"path":        dst,
		"size":        fmt.Sprintf("%.2f KB", float64(fileHeader.Size)/1024),
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
