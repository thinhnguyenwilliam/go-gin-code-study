package v1handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/dto"
	"github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/utils"
)

type ProductHandler struct {
	validate *validator.Validate
}

func NewProductHandler() *ProductHandler {
	v := validator.New()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("alphanumspace", utils.AlphaNumSpace)
		_ = v.RegisterValidation("imgext", utils.ValidateImageExtension)
	}
	return &ProductHandler{validate: v}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	log.Printf("🔍 Request: %s %s from %s", c.Request.Method, c.FullPath(), c.ClientIP())
	log.Printf("📦 Content-Type: %s", c.ContentType())
	log.Printf("🧾 Headers: %+v", c.Request.Header)

	// If needed, read body manually (use with care — affects binding)
	// body, _ := io.ReadAll(c.Request.Body)
	// log.Printf("Raw body: %s", string(body))
	var req dto.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Validation failed",
			"fields": utils.FormatValidationErrors(err),
		})
		return
	}

	// ✅ Set default if field is not provided (still false)
	if req.Display == nil {
		trueVal := true
		req.Display = &trueVal
	}

	// Print the received JSON
	log.Printf("Received product: %+v\n", req)

	//If you want pretty-print JSON, use json.MarshalIndent:
	b, _ := json.MarshalIndent(req, "", "  ")
	log.Println("Request body:", string(b))

	// Auto set created_at timestamp
	now := time.Now().Format("2006-01-02 15:04:05")
	req.CreatedAt = now

	// ✅ Uniqueness check
	if h.ProductNameExists(req.Name) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product name already exists",
			"fields": gin.H{
				"name": "This product name is already in use",
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "New product created",
		"data":    req,
	})
}

// Example with dummy memory check
func (h *ProductHandler) ProductNameExists(name string) bool {
	// Replace with DB logic like:
	// err := db.Where("name = ?", name).First(&product).Error
	existing := []string{"Golang T-shirt", "Python Mug"}

	for _, n := range existing {
		if strings.EqualFold(n, name) {
			return true
		}
	}
	return false
}

// If using a real DB (e.g., GORM + MySQL/PostgreSQL):
// func (h *ProductHandler) ProductNameExists(name string) bool {
// 	var product models.Product
// 	err := h.db.Where("name = ?", name).First(&product).Error
// 	return err == nil // product exists
// }

func (h *ProductHandler) GetProductByLang(c *gin.Context) {
	var uri dto.ProductLangUri

	if err := c.ShouldBindUri(&uri); err != nil {
		// Handle validation error
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Validation failed",
			"fields": utils.FormatValidationErrors(err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"language": uri.Lang,
		"message":  "Products filtered by language: " + uri.Lang,
	})
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	var query dto.ProductQuery

	// Bind query params
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid query parameters",
			"fields": utils.FormatValidationErrors(err),
		})
		return
	}

	// Default value for limit
	if query.Limit == 0 {
		query.Limit = 10
	}

	// Re-validate using custom validator (e.g., alphanumspace, min/max)
	if err := h.validate.Struct(query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Validation failed",
			"fields": utils.FormatValidationErrors(err),
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, gin.H{
		"search":  query.Search,
		"limit":   query.Limit,
		"message": "Success",
		"email":   query.Email,
	})
}
