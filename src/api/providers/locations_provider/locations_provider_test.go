package locations_provider

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

// Possible errors when calling API
// 1. rest client error
// 2. not found
// 3. invalid error interface
// 4. valid response but invalid json format
// 5. valid json format and no error

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	m.Run()
}

func TestGetCountryRestClientError(t *testing.T) {
	rest.FlushMockups()
	apiHeaders := make(http.Header)
	mock := rest.Mock{
		URL:          fmt.Sprintf(UrlGetCountry, "IN"),
		HTTPMethod:   http.MethodGet,
		ReqHeaders:   apiHeaders,
		RespHTTPCode: http.StatusInternalServerError,
		RespBody:     "",
	}
	rest.AddMockups(&mock)

	country, err := GetCountry("IN")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient error when getting country IN", err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	rest.FlushMockups()
	apiHeaders := make(http.Header)
	mock := rest.Mock{
		URL:          fmt.Sprintf(UrlGetCountry, "XYZ"),
		HTTPMethod:   http.MethodGet,
		ReqHeaders:   apiHeaders,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	}
	rest.AddMockups(&mock)

	country, err := GetCountry("XYZ")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	apiHeaders := make(http.Header)
	// mock status as string instead of int
	mock := rest.Mock{
		URL:          fmt.Sprintf(UrlGetCountry, "IN"),
		HTTPMethod:   http.MethodGet,
		ReqHeaders:   apiHeaders,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Country not found","error":"not_found","status":"404","cause":[]}`,
	}
	rest.AddMockups(&mock)

	country, err := GetCountry("IN")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient error when getting country IN", err.Message)
}

func TestGetCountryValidResponseButInvalidFormat(t *testing.T) {
	rest.FlushMockups()
	apiHeaders := make(http.Header)
	// mock `id` in RespBody as int instead of string
	// actual response {"id":"IN","name":"India","time_zone":"GMT+05:30"}
	mock := rest.Mock{
		URL:          fmt.Sprintf(UrlGetCountry, "IN"),
		HTTPMethod:   http.MethodGet,
		ReqHeaders:   apiHeaders,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":123,"name":"India","time_zone":"GMT+05:30"}`,
	}
	rest.AddMockups(&mock)

	country, err := GetCountry("IN")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for IN", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	rest.FlushMockups()
	apiHeaders := make(http.Header)
	mock := rest.Mock{
		URL:          fmt.Sprintf(UrlGetCountry, "IN"),
		HTTPMethod:   http.MethodGet,
		ReqHeaders:   apiHeaders,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"IN","name":"India","locale":"en_US","currency_id":"USD","decimal_separator":".","thousands_separator":",","time_zone":"GMT+05:30","geo_information":{"location":{"latitude":28.6699929,"longitude":77.23000403}},"states":[{"id":"IN-AP","name":"Andhra Pradesh"},{"id":"IN-AR","name":"Arunachal Pradesh"},{"id":"IN-AS","name":"Assam"},{"id":"IN-BR","name":"Bihar"},{"id":"IN-CH","name":"Chandigarh"},{"id":"IN-CT","name":"Chhattisgarh"},{"id":"IN-DN","name":"Dadra And Nagar Haveli"},{"id":"IN-DL","name":"Delhi"},{"id":"IN-GA","name":"Goa"},{"id":"IN-GJ","name":"Gujarat"},{"id":"IN-HR","name":"Haryana"},{"id":"IN-HP","name":"Himachal Pradesh"},{"id":"IN-AN","name":"Andaman And Nicobar Islands"},{"id":"IN-JK","name":"Jammu And Kashmir"},{"id":"IN-JH","name":"Jharkhand"},{"id":"IN-KA","name":"Karnataka"},{"id":"IN-KL","name":"Kerala"},{"id":"IN-MP","name":"Madhya Pradesh"},{"id":"IN-MH","name":"Maharashtra"},{"id":"IN-MN","name":"Manipur"},{"id":"IN-ML","name":"Meghalaya"},{"id":"IN-MZ","name":"Mizoram"},{"id":"IN-NL","name":"Nagaland"},{"id":"IN-OR","name":"Odisha"},{"id":"IN-PY","name":"Puducherry"},{"id":"IN-PB","name":"Punjab"},{"id":"IN-RJ","name":"Rajasthan"},{"id":"IN-TN","name":"Tamil Nadu"},{"id":"IN-TG","name":"Telangana"},{"id":"IN-TR","name":"Tripura"},{"id":"IN-UP","name":"Uttar Pradesh"},{"id":"IN-UT","name":"Uttarakhand"},{"id":"IN-WB","name":"West Bengal"}]}`,
	}
	rest.AddMockups(&mock)

	country, err := GetCountry("IN")
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "IN", country.ID)
	assert.EqualValues(t, "India", country.Name)
	assert.EqualValues(t, "GMT+05:30", country.TimeZone)
	assert.EqualValues(t, 33, len(country.States))
}
