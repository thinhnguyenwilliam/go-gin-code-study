package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/dto"
	"github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/utils"
)

type ProductHandler struct {
	validate *validator.Validate
}

func NewProductHandler() *ProductHandler {
	v := validator.New()
	_ = v.RegisterValidation("alphanumspace", utils.AlphaNumSpace)
	return &ProductHandler{validate: v}
}

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
