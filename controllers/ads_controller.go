package controllers

import (
	"github.com/Nicrii/Advertisement/services"
	"github.com/Nicrii/Advertisement/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAd(c *gin.Context) {
	adId, err := strconv.ParseInt(c.Param("ad_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{Message: "ad_id must be a number", StatusCode: http.StatusBadRequest, Code: "bad_request"}
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	fields := c.QueryArray("fields[]")

	ad, apiErr := services.AdsService.GetAd(adId, fields)
	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	c.JSON(http.StatusOK, ad)
}
