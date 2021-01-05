package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/mmkader85/golang-basic-testing/src/api/domain/locations"
	"github.com/mmkader85/golang-basic-testing/src/api/services"
	"github.com/mmkader85/golang-basic-testing/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc func(countryID string) (*locations.Country, *errors.ApiError)
)

type locationsServiceMock struct{}

func (*locationsServiceMock) GetCountry(countryID string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryID)
}

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {
	getCountryFunc = func(countryID string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{
			Message: "Country not found",
			Error:   "not_found",
			Status:  http.StatusNotFound,
		}
	}
	services.LocationService = &locationsServiceMock{}

	// We dont need to mock Rest API anymore as services.LocationService.GetCountry method,
	// which calls location_provider.GetCountry function,
	// which makes Rest API call, has been mocked.
	// rest.FlushMockups()
	// apiHeaders := make(http.Header)
	// mock := rest.Mock{
	// 	URL:          fmt.Sprintf(UrlGetCountry, "XYZ"),
	// 	HTTPMethod:   http.MethodGet,
	// 	ReqHeaders:   apiHeaders,
	// 	RespHTTPCode: http.StatusNotFound,
	// 	RespBody:     `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	// }
	// rest.AddMockups(&mock)

	testResponse := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(testResponse)
	testContext.Params = gin.Params{
		{Key: "country_id", Value: "XYZ"},
	}

	GetCountry(testContext)
	assert.EqualValues(t, http.StatusNotFound, testResponse.Code)

	var apiErr errors.ApiError
	_ = json.Unmarshal(testResponse.Body.Bytes(), &apiErr)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}

func TestGetCountryNoError(t *testing.T) {
	getCountryFunc = func(countryID string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{
			ID:       "IN",
			Name:     "India",
			TimeZone: "GMT+05:30",
			GeoInformation: locations.GeoInformation{
				Location: locations.GeoLocation{
					Latitude:  12.23,
					Longitude: 14.232,
				}},
			States: []locations.State{
				{
					ID:   "IN-AP",
					Name: "Andhra Pradesh",
				},
				{
					ID:   "IN-TN",
					Name: "Tamil Nadu",
				},
			},
		}, nil
	}
	services.LocationService = &locationsServiceMock{}

	testResponse := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(testResponse)
	testContext.Params = gin.Params{
		{Key: "country_id", Value: "IN"},
	}

	GetCountry(testContext)
	assert.EqualValues(t, http.StatusOK, testResponse.Code)

	var country *locations.Country
	json.Unmarshal(testResponse.Body.Bytes(), &country)
	assert.EqualValues(t, country.Name, "India")
}
