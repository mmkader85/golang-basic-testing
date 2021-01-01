package locations_provider

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/mmkader85/golang-basic-testing/src/api/domain/locations"
	"github.com/mmkader85/golang-basic-testing/src/api/utils/errors"
)

const (
	UrlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

func GetCountry(countryID string) (*locations.Country, *errors.ApiError) {
	response := rest.Get(fmt.Sprintf(UrlGetCountry, countryID))
	if response == nil || response.Response == nil {
		return nil, &errors.ApiError{
			Message: fmt.Sprintf("invalid restclient error when getting country %s", countryID),
			Status:  http.StatusInternalServerError,
		}
	}

	if response.Response.StatusCode > 299 {
		var apiErr errors.ApiError
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
			return nil, &errors.ApiError{
				Message: fmt.Sprintf("invalid restclient error when getting country %s", countryID),
				Status:  http.StatusInternalServerError,
			}
		}
		return nil, &apiErr
	}

	var result locations.Country
	if err := json.Unmarshal(response.Bytes(), &result); err != nil {
		return nil, &errors.ApiError{
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryID),
			Status:  http.StatusInternalServerError,
		}
	}

	return &result, nil
}
