package domain

import "google.golang.org/genproto/googleapis/type/date"

type Ad struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImgURL      []string  `json:"img_url"`
	Price       int       `json:"price"`
	Date        date.Date `json:"date"` ///TODO: проверить/изменить тип
}
