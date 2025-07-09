package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1handler "github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/internal/api/v1/handler"
	"github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/utils"
)

func GetProducts(c *gin.Context) {
	search := c.Query("search")
	limitStr := c.Query("limit")

	// Validate search
	if err := utils.ValidateSearch(search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate limit
	limit, err := utils.ValidateLimit(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Continue with your product logic
	c.JSON(http.StatusOK, gin.H{
		"search":  search,
		"limit":   limit,
		"message": "Products fetched successfully.",
	})
}

func GetProductByIdV1(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Product details for ID " + id,
	})
}

func PostProductV1(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "New product created",
	})
}

func PutProductV1(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated product with ID " + id,
	})
}

func DeleteProductV1(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted product with ID " + id,
	})
}

func GetProductByLang(c *gin.Context) {
	lang := c.Param("lang")

	allowed := map[string]bool{
		"php":    true,
		"golang": true,
		"python": true,
	}

	if !allowed[lang] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid language",
			"message": "Allowed values are: php, golang, python",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"language": lang,
		"message":  "Products filtered by language: " + lang,
	})
}

const (
	userByIDRoute    = "/:id"
	productByIDRoute = "/:id"
)

func main() {
	r := gin.Default()
	userHandler := v1handler.NewUserHandler()

	// Group for version 1
	v1 := r.Group("/api/v1")
	{
		// /api/v1/users group
		users := v1.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/uuid/:uuid", userHandler.GetUsersByUUID)
			users.GET("/slug", userHandler.GetUserWithoutSlug)
			users.GET("/slug/:slug", userHandler.GetUserBySlug)
			users.GET(userByIDRoute, userHandler.GetUserByID)
			users.POST("", userHandler.CreateUser)
			users.PUT(userByIDRoute, userHandler.UpdateUser)
			users.DELETE(userByIDRoute, userHandler.DeleteUser)
		}

		// /api/v1/products group
		products := v1.Group("/products")
		{
			products.GET("", GetProducts)
			products.GET("/category/:lang", GetProductByLang)
			products.GET(productByIDRoute, GetProductByIdV1)
			products.POST("", PostProductV1)
			products.PUT(productByIDRoute, PutProductV1)
			products.DELETE(productByIDRoute, DeleteProductV1)
		}
	}

	r.Run(":8080")
}
