package routers

import (
	"github.com/Vamsi40543/Go_Weather_api/weather_app/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/getTodayWeather", handler.GetTodayWeather)
	router.GET("/weather/history", handler.GetAllWeather)

}
