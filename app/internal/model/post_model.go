package model

type CreatePost struct {
	Description string `form:"description" db:"description"`
	CreatedBy   int    `form:"created_by" db:"created_by"`
}

type Post struct {
	Id          int      `form:"id" db:"id"`
	Description string   `form:"description" db:"description"`
	CreatedBy   int      `form:"created_by" db:"created_by"`
	MainImage   string   `form:"main_image" db:"main_image"`
	Images      []string `form:"images" db:"images"`
}
