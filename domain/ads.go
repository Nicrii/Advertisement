package domain

type Ad struct {
	Id          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImagesURLs  []string `json:"images_urls"`
	Price       int      `json:"price"`
	Date        string   `json:"date"`
}
