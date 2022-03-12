package models

type Post struct {
	Model
	Title       string `json:"title" form:"title"`
	Description string `json:"desc" form:"description"`
	Content     string `json:"content" form:"content"`
	UserID      int    `json:"user_id" form:"user_id"`
	TagID       int    `json:"tag_id" form:"tag_id"`
	User        User
	Tag         Tag
}
