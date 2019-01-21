package errs

import (
	"github.com/mailru/easyjson"
)

//go:generate easyjson

const (
	NotFound = iota
	InvalidFormat
	ConstraintViolation
	Conflict
	NotAuthorized
	PermissionDenied
	InternalError
)

//easyjson:json
type Error struct {
	Type      int    `json:"-"`
	Message   string `json:"message"`
	NestedErr error  `json:"-"`
	JSON      []byte `json:"-"`
}

func NewError(t int, message string) *Error {
	err := &Error{
		Type:      t,
		Message:   message,
		NestedErr: nil,
	}

	err.JSON, _ = easyjson.Marshal(err)

	return err
}

const (
	InternalErrorMessage = "internal service error"
)

func NewServiceError(err error) *Error {
	sErr := NewError(InternalError, InternalErrorMessage)
	sErr.NestedErr = err
	return sErr
}

func (v *Error) Error() string {
	return v.Message
}
