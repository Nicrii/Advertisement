package controllers

import (
	"encoding/json"
	"github.com/Nicrii/Advertisement/services"
	"github.com/Nicrii/Advertisement/utils"
	"net/http"
	"strconv"
)

func GetAd(response http.ResponseWriter, request *http.Request) {
	adId, err := strconv.ParseInt(request.URL.Query().Get("Ad_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{Message: "ad_id must be a number", StatusCode: http.StatusBadRequest, Code: "bad_request"}
		response.WriteHeader(apiErr.StatusCode)
		jsonValue, _ := json.Marshal(apiErr)
		response.Write(jsonValue)
		return
	}

	ad, apiErr := services.GetAd(adId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(ad)

	response.Write(jsonValue)
}
