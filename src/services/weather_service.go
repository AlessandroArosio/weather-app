package services2

import (
	"encoding/json"
	"fmt"
	"github.com/alessandroarosio/weather-workshop/src/model/errors"
	"github.com/alessandroarosio/weather-workshop/src/model/weather"
	"github.com/go-resty/resty"
	"net/http"
	"os"
	"strings"
)

const (
	baseUrl = "https://api.weatherbit.io/v2.0/current"
	apiKey = "api_key"
)

var (
	WeatherService weatherServiceInterface = &weatherService{}
	key = os.Getenv(apiKey)
)

type weatherService struct{}

type weatherServiceInterface interface {
	GetWeather(city string, country string) (*weather.WeatherResponse, *errors.AppError)
}

func (ws *weatherService) GetWeather(city string, country string) (*weather.WeatherResponse, *errors.AppError) {
	city = strings.TrimSpace(city)
	country = strings.ToUpper(strings.TrimSpace(country))

	endpoint := fmt.Sprintf("?city=%s&country=%s&key=%s", city, country, key)

	client := resty.New()
	response, _ := client.R().Get(baseUrl + endpoint)

	var weatherResponse weather.WeatherResponse
	if err := json.Unmarshal(response.Body(), &weatherResponse); err != nil {
		msg := "error mapping to struct"
		return nil, errors.NewGenericError(msg, http.StatusInternalServerError)
	}

	return &weatherResponse, nil
}
