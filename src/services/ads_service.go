package services

import (
	"github.com/Nicrii/Advertisement/domain"
	"github.com/Nicrii/Advertisement/utils"
	"net/http"
	"strconv"
)

type adsService struct {
}

var (
	AdsService adsService
)

func (a *adsService) Get(adId int64, fields []string) (*domain.GetResponse, *utils.ApplicationError) {
	if adId <= 0 {
		return nil, &utils.ApplicationError{Message: "Ad_id must be more than 0", Code: http.StatusBadRequest}
	}
	result, err := domain.GetAd(adId, fields)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adsService) Create(ad domain.Ad) (int64, int) {
	if status := ad.Validate(); status != http.StatusOK {
		return 0, status
	}

	return ad.Save()
}

func (a *adsService) GetList(page, sortBy, sortDirection string) (*[]domain.GetResponse, *utils.ApplicationError) {

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	if pageInt <= 0 {
		pageInt = 1
	}

	isCorrect := false
	for _, field := range []string{"price", "date"} {
		if sortBy == field {
			isCorrect = true
			break
		}
	}
	if !isCorrect {
		return nil, &utils.ApplicationError{Message: "Unsupported sortBy value", Code: http.StatusBadRequest}
	}

	isCorrect = false
	for _, field := range []string{"asc", "desc"} {
		if sortDirection == field {
			isCorrect = true
			break
		}
	}
	if !isCorrect {
		return nil, &utils.ApplicationError{Message: "Unsupported sortDirection value", Code: http.StatusBadRequest}
	}

	result, apiErr := domain.GetListOfAds(pageInt, sortBy, sortDirection)
	if apiErr != nil {
		return nil, apiErr
	}

	return result, nil
}
