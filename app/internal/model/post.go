package model

type Post struct {
	Description string   `json:"description"`
	MainPhoto   string   `json:"main_photo"`
	Photos      string   `json:"photos"`
	Views       int      `json:"views"`
	Likes       int      `json:"likes"`
	Comments    int      `json:"comments"`
	Hashtags    []string `json:"hashtags"`
}
