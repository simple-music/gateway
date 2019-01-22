package models

//go:generate easyjson

//easyjson:json
type SubscriptionsStatus struct {
	User             string `json:"user"`
	NumSubscribers   int64  `json:"numSubscribers"`
	NumSubscriptions int64  `json:"numSubscriptions"`
}

//easyjson:json
type UserFull struct {
	ID                 string   `json:"id"`
	Username           string   `json:"username"`
	Email              string   `json:"email"`
	DateOfBirth        *string  `json:"dateOfBirth"`
	MusicalInstruments []string `json:"musicalInstruments"`
	NumSubscribers     int64    `json:"numSubscribers"`
	NumSubscriptions   int64    `json:"numSubscriptions"`
}

func (v *UserFull) From(musician *Musician, status *SubscriptionsStatus) {
	v.ID = musician.ID
	v.Username = musician.Nickname
	v.Email = musician.Email
	v.DateOfBirth = musician.DateOfBirth
	v.MusicalInstruments = musician.MusicalInstruments
	v.NumSubscribers = status.NumSubscribers
	v.NumSubscriptions = status.NumSubscriptions
}

//easyjson:json
type UsersPage []string
