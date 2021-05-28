package controllers

import (
	"fmt"
	"github.com/Nicrii/Advertisement/domain"
	"github.com/Nicrii/Advertisement/services"
	"github.com/Nicrii/Advertisement/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Get(c *gin.Context) {
	adId, err := strconv.ParseInt(c.Param("ad_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{Message: "ad_id must be a number", StatusCode: http.StatusBadRequest, Code: "bad_request"}
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	fields := c.QueryArray("fields[]")

	ad, apiErr := services.AdsService.Get(adId, fields)
	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	c.JSON(http.StatusOK, ad)
}

func Create(c *gin.Context) {
	var ad domain.Ad
	if err := c.ShouldBindJSON(&ad); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":          0,
			"status_code": http.StatusBadRequest,
		})
		fmt.Println(err)
		return

	}

	id, statusCode := services.AdsService.Create(ad)
	c.JSON(statusCode, gin.H{
		"id":          id,
		"status_code": statusCode,
	})
}
