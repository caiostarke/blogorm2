package models

type Tag struct {
	Model
	TagName     string `json:"nome"`
	Description string `json:"description"`
}
