package domain

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
