package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	router := gin.Default()

	// GET /ping -> "pong"
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// GET /hello/:name?uid=123
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")              // from path
		uid := c.DefaultQuery("uid", "0000") // from query param

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name,
			"uid":     uid,
		})
	})

	// POST /login with JSON
	router.POST("/login", func(c *gin.Context) {
		var json struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   "logged in",
			"username": json.Username,
		})
	})

	router.GET("/users", func(c *gin.Context) {
		users := []string{"Alice", "Bob", "Charlie"}

		c.JSON(200, gin.H{
			"data": users,
		})
	})

	router.GET("/user/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id") // Get path param from URL

		// Simulate fetching user data (you can replace this with DB logic)
		user := gin.H{
			"id":   userID,
			"name": "User " + userID,
		}

		c.JSON(200, gin.H{
			"data": user,
		})
	})

	router.GET("/products", func(c *gin.Context) {
		// products := []gin.H{
		// 	{"id": 1, "name": "Laptop", "price": 1000},
		// 	{"id": 2, "name": "Phone", "price": 500},
		// 	{"id": 3, "name": "Headphones", "price": 100},
		// }

		products := []Product{
			{ID: 1, Name: "Laptop", Price: 999},
			{ID: 2, Name: "Phone", Price: 500},
			{ID: 3, Name: "Headphones", Price: 100},
		}

		c.JSON(200, gin.H{
			"data": products,
		})
	})

	router.GET("/product/:product_name", func(c *gin.Context) {
		productName := c.Param("product_name")

		// Query parameters
		priceStr := c.Query("price")
		color := c.Query("color")

		//Note: Query() always returns a string. You can convert it to an integer like this:
		priceInt, err := strconv.Atoi(priceStr)
		if err != nil {
			// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
			// return

			// Default value or error response if not a valid integer
			priceInt = 0
		}

		product := gin.H{
			"name":  productName,
			"price": priceInt,
			"color": color,
			"stock": 20,
		}

		c.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	})

	router.Run(":8080")
}
