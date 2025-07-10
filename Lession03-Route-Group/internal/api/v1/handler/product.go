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

func (h *ProductHandler) GetProducts(c *gin.Context) {
	var query dto.ProductQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	if query.Limit == 0 {
		query.Limit = 10
	}

	if err := h.validate.Struct(query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"search":  query.Search,
		"limit":   query.Limit,
		"message": "Success",
	})
}
