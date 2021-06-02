package controllers

import (
	"github.com/Nicrii/Advertisement/domain"
	"github.com/Nicrii/Advertisement/services"
	"github.com/Nicrii/Advertisement/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Get(ctx *gin.Context) {
	adId, err := strconv.ParseInt(ctx.Param("ad_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{Message: "ad_id must be a number", Code: http.StatusBadRequest}
		ctx.JSON(apiErr.Code, apiErr)

		return
	}

	fields := ctx.QueryArray("fields[]")

	ad, apiErr := services.AdsService.Get(adId, fields)

	if apiErr != nil {
		ctx.JSON(apiErr.Code, apiErr)
		return
	}

	ctx.JSON(http.StatusOK, ad)
}

func Create(ctx *gin.Context) {
	var ad domain.Ad

	if err := ctx.ShouldBindJSON(&ad); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"id":          0,
			"status_code": http.StatusBadRequest,
		})

		return
	}

	id, statusCode := services.AdsService.Create(ad)

	ctx.JSON(statusCode, gin.H{
		"id":          id,
		"status_code": statusCode,
	})
}

func GetList(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	sortBy := ctx.DefaultQuery("sortBy", "price")
	sortDirection := ctx.DefaultQuery("sortDirection", "asc")

	ads, apiErr := services.AdsService.GetList(page, sortBy, sortDirection)
	if apiErr != nil {
		ctx.JSON(apiErr.Code, apiErr)
		return
	}

	ctx.JSON(http.StatusOK, ads)
}
