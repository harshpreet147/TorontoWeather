// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database connection
var db *gorm.DB

func init() {
	// Open database connection using gorm with SQLite
	var err error
	db, err = gorm.Open(sqlite.Open("file:toronto_time.db?cache=shared&_loc=auto"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto Migrate the time_log table
	db.AutoMigrate(&TimeLog{})
}

// TimeLog struct for the time_log table
type TimeLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

// API endpoint to get all times from the time_log table
func getAllTimes(c *gin.Context) {
	var timeLogs []TimeLog

	// Retrieve all entries from the time_log table
	result := db.Find(&timeLogs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve time logs"})
		return
	}

	// Respond with the retrieved time logs in JSON format
	c.JSON(http.StatusOK, timeLogs)
}

// API endpoint to get current time in Toronto
func getCurrentTime(c *gin.Context) {
	// Get current time in Toronto
	torontoTimeZone, err := time.LoadLocation("America/Toronto")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load Toronto timezone"})
		return
	}
	currentTime := time.Now().In(torontoTimeZone)

	// Create a TimeLog instance
	timeLog := TimeLog{
		Timestamp: currentTime,
	}

	// Insert current time into the database using gorm
	result := db.Create(&timeLog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log timestamp"})
		return
	}

	// Respond with the current time in JSON format
	c.JSON(http.StatusOK, gin.H{"current_time": currentTime.Format(time.RFC3339)})
}

func main() {
	// Set up Gin router
	router := gin.Default()

	// Define API routes
	router.GET("/current-time", getCurrentTime)
	router.GET("/time", getCurrentTime)
	router.GET("/all-times", getAllTimes)

	// Run the server
	port := 7575
	addr := fmt.Sprintf(":%d", port)
	err := router.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}
