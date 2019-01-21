package models

//go:generate easyjson

const (
	UsernamePattern   = `^\w+$`
	EmailPattern      = `^.+@.+$`
	FullNamePattern   = `^((\w)+ ?)+$`
	DatePattern       = `^\d{4}-\d{2}-\d{2}$`
	PasswordMinLength = 8
)

//easyjson:json
type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

//easyjson:json
type NewMusician struct {
	Nickname           string   `json:"nickname"`
	Email              string   `json:"email"`
	FullName           string   `json:"fullName"`
	DateOfBirth        *string  `json:"dateOfBirth"`
	MusicalInstruments []string `json:"musicalInstruments"`
}

func (v *NewMusician) From(newUser *NewUser) {
	v.Nickname = newUser.Username
	v.Email = newUser.Email
	v.FullName = newUser.FullName
}

//easyjson:json
type Musician struct {
	ID                 string   `json:"id"`
	Nickname           string   `json:"nickname"`
	Email              string   `json:"email"`
	FullName           string   `json:"fullName"`
	DateOfBirth        *string  `json:"dateOfBirth"`
	MusicalInstruments []string `json:"musicalInstruments"`
}

//easyjson:json
type MusicianUpdate struct {
	Email              *string  `json:"email"`
	FullName           *string  `json:"fullName"`
	DateOfBirth        *string  `json:"dateOfBirth"`
	MusicalInstruments []string `json:"musicalInstruments"`
}

//easyjson:json
type NewCredentials struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (v *NewCredentials) From(newUser *NewUser, musician *Musician) {
	v.UserID = musician.ID
	v.Username = musician.Nickname
	v.Password = newUser.Password
}

//easyjson:json
type AuthCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
