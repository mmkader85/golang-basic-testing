package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mmkader85/golang-basic-testing/src/api/services"
	"github.com/mmkader85/golang-basic-testing/src/api/utils/errors"
)

func GetCountry(c *gin.Context) {
	countryID := c.Param("country_id")
	if strings.TrimSpace(countryID) == "" {
		c.JSON(http.StatusBadRequest, &errors.ApiError{
			Message: "country_id param is missing",
			Error:   "bad_request",
			Status:  http.StatusBadRequest,
		})
		return
	}

	country, err := services.LocationService.GetCountry(countryID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, country)
}
