package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mmkader85/golang-basic-testing/src/api/controllers"
)

var router = gin.Default()

func mapURLs() { router.GET("/location/country/:country_id", controllers.GetCountry) }
