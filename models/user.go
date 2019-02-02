package models

import (
	"github.com/simple-music/gateway/errs"
	"regexp"
)

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

type NewUserValidator struct {
	usernameRegexp *regexp.Regexp
	emailRegexp    *regexp.Regexp
	fullNameRegexp *regexp.Regexp

	usernameErr *errs.Error
	passwordErr *errs.Error
	emailErr    *errs.Error
	fullNameErr *errs.Error
}

func NewNewUserValidator() *NewUserValidator {
	return &NewUserValidator{
		usernameRegexp: regexp.MustCompile(UsernamePattern),
		emailRegexp:    regexp.MustCompile(EmailPattern),
		fullNameRegexp: regexp.MustCompile(FullNamePattern),

		usernameErr: errs.NewError(errs.InvalidFormat, "invalid username"),
		passwordErr: errs.NewError(errs.InvalidFormat, "invalid password"),
		emailErr:    errs.NewError(errs.InvalidFormat, "invalid email"),
		fullNameErr: errs.NewError(errs.InvalidFormat, "invalid full name"),
	}
}

func (v *NewUserValidator) Validate(user *NewUser) *errs.Error {
	if !v.usernameRegexp.MatchString(user.Username) {
		return v.usernameErr
	}
	if len(user.Password) < PasswordMinLength {
		return v.passwordErr
	}
	if !v.emailRegexp.MatchString(user.Email) {
		return v.emailErr
	}
	if !v.fullNameRegexp.MatchString(user.FullName) {
		return v.fullNameErr
	}
	return nil
}

type MusicianUpdateValidator struct {
	emailRegexp    *regexp.Regexp
	fullNameRegexp *regexp.Regexp
	dateRegexp     *regexp.Regexp

	emailErr       *errs.Error
	fullNameErr    *errs.Error
	dateOfBirthErr *errs.Error
}

func NewMusicianUpdateValidator() *MusicianUpdateValidator {
	return &MusicianUpdateValidator{
		emailRegexp:    regexp.MustCompile(EmailPattern),
		fullNameRegexp: regexp.MustCompile(FullNamePattern),
		dateRegexp:     regexp.MustCompile(DatePattern),

		emailErr:       errs.NewError(errs.InvalidFormat, "invalid email"),
		fullNameErr:    errs.NewError(errs.InvalidFormat, "invalid full name"),
		dateOfBirthErr: errs.NewError(errs.InvalidFormat, "invalid date of birth"),
	}
}

func (v *MusicianUpdateValidator) Validate(update *MusicianUpdate) *errs.Error {
	if update.Email != nil && !v.emailRegexp.MatchString(*update.Email) {
		return v.emailErr
	}
	if update.FullName != nil && !v.fullNameRegexp.MatchString(*update.FullName) {
		return v.fullNameErr
	}
	if update.DateOfBirth != nil && !v.dateRegexp.MatchString(*update.DateOfBirth) {
		return v.dateOfBirthErr
	}
	return nil
}
