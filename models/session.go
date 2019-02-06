package models

//go:generate easyjson

//easyjson:json
type Session struct {
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
	UserID       string `json:"userId"`
}
