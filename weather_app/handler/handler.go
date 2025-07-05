package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Vamsi40543/Go_Weather_api/weather_app/config"
	"github.com/Vamsi40543/Go_Weather_api/weather_app/database"
	"github.com/Vamsi40543/Go_Weather_api/weather_app/models"
	"github.com/gin-gonic/gin"
)

func GetTodayWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City name is required"})
		return
	}

	apiKey := config.App.OpenWeatherMapApiKey
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("🌐 Raw API Response:\n", string(body))
	fmt.Println("🔁 HTTP Status Code:", resp.StatusCode)

	var data models.OpenWeatherResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("❌ JSON Unmarshal Error:", err)
		fmt.Println("🔎 Raw body again:\n", string(body))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse weather data"})
		return
	}
	fmt.Printf("✅ Parsed OpenWeatherResponse: %+v\n", data)

	now := time.Now()
	weather := models.Weather{
		Date:        now.Format("2006-01-02"),
		Time:        now.Format("15:04:05"),
		City:        data.Name,
		Latitude:    data.Coord.Lat,
		Longitude:   data.Coord.Lon,
		Temperature: data.Main.Temp,
		Humidity:    data.Main.Humidity,
	}

	// Debug log before DB save
	fmt.Printf("✅ Weather to insert: %+v\n", weather)

	if err := database.DB.Create(&weather).Error; err != nil {
		fmt.Println("❌ DB Save Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save weather to DB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "✅ Weather fetched and saved successfully",
		"data":    weather,
	})
}

func GetAllWeather(c *gin.Context) {
	var weathers []models.Weather
	if err := database.DB.Find(&weathers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weathers"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": weathers})
}
