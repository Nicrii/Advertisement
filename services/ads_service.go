package services

import (
	"github.com/Nicrii/Advertisement/domain"
	"github.com/Nicrii/Advertisement/utils"
)

func GetAd(adId int64) (*domain.Ad, *utils.ApplicationError) {
	return domain.GetAd(adId)
}
