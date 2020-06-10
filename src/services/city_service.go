package services2

import (
	"github.com/alessandroarosio/weather-workshop/src/model/city"
	"github.com/alessandroarosio/weather-workshop/src/model/errors"
)

var (
	CityService cityServiceInterface = &cityService{}
)

type cityService struct {}

type cityServiceInterface interface {
	SaveCity(city city.City) (*city.City, *errors.AppError)
}

func (cs *cityService) SaveCity(city city.City) (*city.City, *errors.AppError) {
	if err := city.Validate(); err != nil {
		return nil, err
	}
	if err := city.Save(); err != nil {
		return nil, err
	}

	return &city, nil
}
