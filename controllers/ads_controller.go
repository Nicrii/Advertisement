package controllers

import (
	"github.com/Nicrii/Advertisement/services"
	"github.com/Nicrii/Advertisement/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ResultGet struct {
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	MainImage   string   `json:"main_image"`
	Description string   `json:"description,omitempty"`
	ImagesURLs  []string `json:"images_urls,omitempty"`
}

func GetAd(c *gin.Context) {
	var result ResultGet
	adId, err := strconv.ParseInt(c.Param("ad_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{Message: "ad_id must be a number", StatusCode: http.StatusBadRequest, Code: "bad_request"}
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	fields := c.QueryArray("fields[]")

	ad, apiErr := services.AdsService.GetAd(adId)
	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	result.Name = ad.Name
	result.Price = ad.Price
	result.MainImage = ad.ImagesURLs[0]

	for _, field := range fields {
		switch field {
		case "description":
			result.Description = ad.Description
		case "images_urls":
			result.ImagesURLs = ad.ImagesURLs
		}
	}

	c.JSON(http.StatusOK, result)
}
