package city

import (
	"github.com/alessandroarosio/weather-workshop/src/datasource/mysql"
	"github.com/alessandroarosio/weather-workshop/src/model/errors"
	"github.com/apex/log"
	"net/http"
)

const (
	queryInsertCity = "INSERT INTO city (city_id, city_name, state_code, country_code, country_full, lat, lon) VALUES (?, ?, ?, ?, ?, ?, ?);"
)

func (city *City) Save() *errors.AppError {
	stmt, err := mysql.Client.Prepare(queryInsertCity)
	if err != nil {
		errorMessage := "error when preparing SaveCity statement"
		log.Error(errorMessage)
		return errors.NewGenericError(errorMessage, http.StatusInternalServerError)
	}
	defer stmt.Close()

	_, saveErr := stmt.Exec(city.Id, city.CityName, city.StateCode, city.CountryCode, city.Country, city.Lat, city.Lon)
	if saveErr != nil {
		errorMessage := "error when saving city to DB"
		log.Error(errorMessage)
		return errors.NewGenericError(errorMessage, http.StatusInternalServerError)
	}
	return nil
}