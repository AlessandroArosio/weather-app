package app

import "github.com/alessandroarosio/weather-workshop/src/controllers"

func mapUrls() {
	router.GET("/weather/:city/:country", controllers.GetWeather)
}