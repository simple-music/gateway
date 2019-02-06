package models

//go:generate easyjson

//easyjson:json
type AuthCode struct {
	ClientID     string `json:"-"`
	ClientSecret string `json:"-"`
	AuthCode     string `json:"authCode"`
}
