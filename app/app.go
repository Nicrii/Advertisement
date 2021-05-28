package app

import (
	"github.com/Nicrii/Advertisement/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/GetAd", controllers.GetAd)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
