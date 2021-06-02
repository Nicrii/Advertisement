package domain

import (
	"fmt"
	"github.com/Nicrii/Advertisement/datasources/postgresql/ads_db"
	"github.com/Nicrii/Advertisement/utils"
	"github.com/lib/pq"
	"net/http"
)

const (
	queryGetAd         = "SELECT id,name,description,images_urls,price,date FROM advertisement WHERE id=$1;"
	queryInsertAd      = "INSERT INTO advertisement(name,description,images_urls,price,date) VALUES($1, $2, $3,$4,now()) RETURNING id;"
	queryGetListSelect = "SELECT name,images_urls[1],price FROM advertisement ORDER BY"
	queryGetListLimit  = "LIMIT 10 OFFSET ($1-1)*10"
)

func GetAd(id int64, fields []string) (*GetResponse, *utils.ApplicationError) {
	var ad Ad
	var result GetResponse

	stmt, err := ads_db.Client.Prepare(queryGetAd)
	if err != nil {
		return nil, &utils.ApplicationError{Message: fmt.Sprintf("error when trying to get ad %d", id), Code: http.StatusInternalServerError}
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(&ad.Id, &ad.Name, &ad.Description, pq.Array(&ad.ImagesURLs), &ad.Price, &ad.Date); err != nil {
		return nil, &utils.ApplicationError{Message: fmt.Sprintf("error when trying to get ad %d", id), Code: http.StatusNotFound}
	}
	result.Name = ad.Name
	result.Price = ad.Price

	for _, field := range fields {
		switch field {
		case "description":
			result.Description = ad.Description
			break
		case "images_urls":
			result.ImagesURLs = ad.ImagesURLs
			break
		}
	}

	return &result, nil
}

func (ad *Ad) Save() (int64, int) {
	var id int64

	stmt, err := ads_db.Client.Prepare(queryInsertAd)
	if err != nil {
		return 0, http.StatusInternalServerError
	}
	defer stmt.Close()

	insertResult := stmt.QueryRow(ad.Name, ad.Description, pq.Array(ad.ImagesURLs), ad.Price)

	err = insertResult.Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, http.StatusInternalServerError
	}

	return id, http.StatusOK
}

func GetListOfAds(page int, sortBy, sortDirection string) (*[]GetResponse, *utils.ApplicationError) {
	adsList := make([]GetResponse, 0, 10)

	stmt, err := ads_db.Client.Prepare(queryGetListSelect + " " + sortBy + " " + sortDirection + " " + queryGetListLimit)
	if err != nil {
		return nil, &utils.ApplicationError{Message: "Error when trying to get list of ads", Code: http.StatusInternalServerError}
	}
	defer stmt.Close()

	rows, err := stmt.Query(page)
	if err != nil {
		return nil, &utils.ApplicationError{Message: "Error when trying to get list of ads", Code: http.StatusInternalServerError}
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		adsList = append(adsList, GetResponse{})
		if err := rows.Scan(&adsList[i].Name, &adsList[i].MainImage, &adsList[i].Price); err != nil {
			return nil, &utils.ApplicationError{Message: fmt.Sprintf("error when trying to get ad %s", err.Error()), Code: http.StatusNotFound}
		}
	}

	return &adsList, nil
}
