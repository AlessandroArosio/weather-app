package controllers

import (
	services2 "github.com/alessandroarosio/weather-workshop/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWeather(c *gin.Context) {
	city := c.Param("city")
	country := c.Param("country")

	result, err := services2.WeatherService.GetWeather(city, country)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
