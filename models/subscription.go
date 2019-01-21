package models

//go:generate easyjson

//easyjson:json
type SubscriptionsStatus struct {
	User             string `json:"user"`
	NumSubscribers   int64  `json:"numSubscribers"`
	NumSubscriptions int64  `json:"numSubscriptions"`
}

//easyjson:json
type UsersPage []string
