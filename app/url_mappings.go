package app

import (
	"github.com/Nicrii/Advertisement/controllers"
)

func mapUrls() {
	router.GET("/GetAd/:ad_id", controllers.GetAd)
}
