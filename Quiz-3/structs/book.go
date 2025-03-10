package structs

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       string    `json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `json:"thickness"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	CategoryID  int       `json:"category_id"`
}
