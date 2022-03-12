package models

type PostForm struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}
