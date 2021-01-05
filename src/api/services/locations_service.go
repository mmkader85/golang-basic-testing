package services

import (
	"fmt"

	"github.com/mmkader85/golang-basic-testing/src/api/domain/locations"
	"github.com/mmkader85/golang-basic-testing/src/api/providers/locations_provider"
	"github.com/mmkader85/golang-basic-testing/src/api/utils/errors"
)

type locationsServiceStruct struct{}

type locationServiceInterface interface {
	GetCountry(countryID string) (*locations.Country, *errors.ApiError)
}

var LocationService locationServiceInterface = &locationsServiceStruct{}

func (l locationsServiceStruct) GetCountry(countryID string) (*locations.Country, *errors.ApiError) {
	fmt.Println("inside service")
	return locations_provider.GetCountry(countryID)
}

// func GetCountry(countryID string) (*locations.Country, *errors.ApiError) {
// 	return locations_provider.GetCountry(countryID)
// }
