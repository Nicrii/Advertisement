package services

import (
	"github.com/Nicrii/Advertisement/domain"
	"github.com/Nicrii/Advertisement/utils"
)

type adsService struct {
}

var (
	AdsService adsService
)

func (a *adsService) GetAd(adId int64, fields []string) (*domain.GetResponse, *utils.ApplicationError) {
	//result := &domain.GetResponse{Id: adId}
	result, err := domain.GetAd(adId, fields)
	if err != nil {
		return nil, err
	}
	return result, nil
}
