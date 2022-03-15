package models

type User struct {
	Model
	FirstName     string `json:"first_name"`
	Email         string `json:"email"`
	LastName      string `json:"last_name"`
	Job           string `json:"job"`
	ProfilePic    string `json:"profile_pic"`
	BackgroundPic string `json:"background_pic"`
	IsAdmin       bool   `json:"is_admin"`
	Password      []byte `json:"-"`
}
