package config

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Appconfig struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

var App Appconfig

func LoadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)

	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&App)
	if err != nil {
		log.Fatalf("Failed to decode config: %v", err)
	}

}

var LogFile *os.File

func LoggerMiddleware() gin.HandlerFunc {
	var err error
	LogFile, err = os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Set global logger output to the file
	log.SetOutput(LogFile)

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		log.Printf("[%s] %s %s %d %s",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}
