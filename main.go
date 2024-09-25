package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/surajNirala/hotel_services/app/config"
	"github.com/surajNirala/hotel_services/app/databases"
	"github.com/surajNirala/hotel_services/app/models"
	"github.com/surajNirala/hotel_services/app/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Gin mode based on environment
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	// Get APP_PORT and APP_URL from environment variables
	app_port := os.Getenv("APP_PORT")
	app_url := os.Getenv("APP_URL")
	if app_port == "" {
		app_port = "8000"
	}
	if app_url == "" {
		app_url = "localhost"
	}

	// Initialize the database and handle errors
	databases.DatabaseUp()
	route := gin.Default()
	CreateAdminUser()
	// Use custom recovery middleware
	route.Use(CustomRecovery)
	route.Use(NoCacheStore)
	routes.ApiRoutes(route)
	fmt.Println(app_url + ":" + app_port)
	// Print all registered routes
	route.Run(":" + app_port)
}

// CustomRecovery is a middleware function that recovers from any panics and writes a 500 error.
func CustomRecovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// Log the panic details along with stack trace
			log.Printf("Panic recovered: %s\n", err)
			log.Printf("Stack trace: %s\n", debug.Stack())

			// Return a 500 Internal Server Error response
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":      500,
				"message":     "Something Internal Server Error.",
				"panic_error": fmt.Sprintf("An unexpected error occurred: %s", err),
			})
			c.Abort()
		}
	}()
	c.Next()
}

func CreateAdminUser() error {
	DB := config.DB
	customSlice := models.Hotel{
		Name:   "Taj Hotel",
		UserID: 1,
	}
	DB.Create(&customSlice)
	customSlice1 := models.Hotel{
		Name:   "Maurya Hotel",
		UserID: 1,
	}
	DB.Create(&customSlice1)
	customSlice2 := models.Hotel{
		Name:   "Ashoka Hotel",
		UserID: 1,
	}
	DB.Create(&customSlice2)
	return nil
}

func NoCacheStore(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
	}
	c.Next()
}
