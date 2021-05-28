package app

import (
	"github.com/Nicrii/Advertisement/controllers"
)

func mapUrls() {
	router.GET("/ad/:ad_id", controllers.Get)
	router.POST("/ad", controllers.Create)
	//router.GET("/ads",controllers.Create)
}
