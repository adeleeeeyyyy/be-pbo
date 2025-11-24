package dto

type BooksCreateRequest struct {
	Title       string `form:"title"`
	Author      string `form:"author"`
	Publisher   string `form:"publisher"`
	Category    string `form:"category"`
	Stock       int    `form:"stock"`
	Description string `form:"description"`
	Image       string `form:"image"`
	Status      string `form:"status"`
}

type BooksUpdateRequest struct {
	Title       *string `form:"title"`
	Author      *string `form:"author"`
	Publisher   *string `form:"publisher"`
	Category    *string `form:"category"`
	Stock       *int    `form:"stock"`
	Description *string `form:"description"`
	Image       *string `form:"image"`
	Status      *string `form:"status"`
}