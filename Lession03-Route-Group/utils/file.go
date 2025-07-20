package utils

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// GenerateUUIDFileName returns a new file name based on UUID and original extension
func GenerateUUIDFileName(originalName string) string {
	ext := filepath.Ext(originalName)
	return uuid.New().String() + ext
}

// ValidateImageMIME checks whether the uploaded file is a real image
func ValidateImageMIME(file multipart.File) error {
	// Read first 512 bytes to detect content type
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return fmt.Errorf("failed to read file content: %w", err)
	}

	// Reset read pointer (very important)
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to reset file pointer: %w", err)
	}

	contentType := http.DetectContentType(buffer)
	allowedMIMEs := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	if !allowedMIMEs[contentType] {
		return fmt.Errorf("invalid file content type: %s", contentType)
	}
	return nil
}

// AllowedExtensions defines supported file extensions.
var AllowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".mp4":  true,
}

// ValidateFileExtension checks if the file has an allowed extension.
func ValidateFileExtension(filename string, allowed map[string]bool) error {
	ext := strings.ToLower(filepath.Ext(filename))
	if !allowed[ext] {
		return fmt.Errorf("unsupported file type '%s'. Allowed types: %s", ext, allowedList(allowed))
	}
	return nil
}

// ValidateFileSize checks if the file size is under maxBytes
func ValidateFileSize(size int64, maxBytes int64) error {
	if size > maxBytes {
		return fmt.Errorf("file is too large (%.2f MB). Max allowed is %.2f MB",
			float64(size)/(1024*1024), float64(maxBytes)/(1024*1024))
	}
	return nil
}

// Helper: Converts allowed extensions to comma-separated string
func allowedList(allowed map[string]bool) string {
	exts := make([]string, 0, len(allowed))
	for ext := range allowed {
		exts = append(exts, strings.TrimPrefix(ext, "."))
	}
	return strings.Join(exts, ", ")
}
