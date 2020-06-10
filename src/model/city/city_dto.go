package city

import (
	"github.com/alessandroarosio/weather-workshop/src/model/errors"
	"net/http"
	"strings"
)

type City struct {
	Id          int64  `json:"city_id"`
	CityName    string `json:"city_name"`
	StateCode   string `json:"state_code"`
	CountryCode string `json:"country_code"`
	Country     string `json:"country_full"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}

func (city *City) Validate() *errors.AppError {
	if strings.TrimSpace(city.CityName) == "" {
		return errors.NewGenericError("city name should not be empty", http.StatusBadRequest)
	}
	return nil
}
