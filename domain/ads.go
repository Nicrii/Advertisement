package domain

import (
	"net/http"
	"strings"
)

type Ad struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImagesURLs  []string `json:"images_urls"`
	Price       int      `json:"price"`
	Date        string   `json:"date"`
}
type GetResponse struct {
	Name        string   `json:"name"`
	Price       int      `json:"price"`
	MainImage   string   `json:"main_image"`
	Description string   `json:"description,omitempty"`
	ImagesURLs  []string `json:"images_urls,omitempty"`
}

func (ad *Ad) Validate() int {
	ad.Name = strings.TrimSpace(ad.Name)
	if len(ad.Name) <= 0 {
		return http.StatusBadRequest
	}
	ad.Description = strings.TrimSpace(ad.Description)
	for i, _ := range ad.ImagesURLs {
		ad.ImagesURLs[i] = strings.TrimSpace(ad.ImagesURLs[i])
		if i > 2 {
			return http.StatusBadRequest
		}
	}
	if ad.Price < 0 {
		return http.StatusBadRequest
	}

	return http.StatusOK
}
