package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1handler "github.com/thinhnguyen-com/CodeWithTuan/Lession03-Route-Group/internal/api/v1/handler"
)

func GetProductByIdV1(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Product details for ID " + id,
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

const (
	userByIDRoute    = "/:id"
	productByIDRoute = "/:id"
)

func main() {
	r := gin.Default()
	userHandler := v1handler.NewUserHandler()
	productHandler := v1handler.NewProductHandler()
	categoryHandler := v1handler.NewCategoryHandler()

	// Serve files from "uploads" folder under /static/ path
	r.Static("/api/static/categories", "./uploads/categories")

	// Group for version 1
	v1 := r.Group("/api/v1")
	{
		// /api/v1/users group
		users := v1.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/uuid/:uuid", userHandler.GetUserByUUID)
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
			products.GET("", productHandler.GetProducts)
			products.GET("/category/:lang", productHandler.GetProductByLang)
			products.GET(productByIDRoute, GetProductByIdV1)
			products.POST("", productHandler.CreateProduct)
			products.PUT(productByIDRoute, PutProductV1)
			products.DELETE(productByIDRoute, DeleteProductV1)
		}

		categories := v1.Group("/categories")
		{
			categories.POST("", categoryHandler.CreateCategory)
			categories.POST("/upload", categoryHandler.UploadCategoryImage)
			categories.POST("/upload-multiple", categoryHandler.UploadMultipleCategoryImages)
		}
	}

	r.Run(":8080")
}
