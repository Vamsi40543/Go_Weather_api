package main

import (
	"github.com/Vamsi40543/Go_Weather_api/weather_app/config"
	"github.com/Vamsi40543/Go_Weather_api/weather_app/database"
	"github.com/Vamsi40543/Go_Weather_api/weather_app/routers"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig(".apiconfig")
	database.Connect()

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(config.LoggerMiddleware())

	routers.RegisterRoutes(r)

	r.Run(":8080")
}
