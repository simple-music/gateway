package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/simple-music/gateway/config"
	"github.com/simple-music/gateway/consts"
	"github.com/simple-music/gateway/errs"
	"time"
)

type TokenManager struct {
	key []byte

	invalidTokenErr  *errs.Error
	notAuthorizedErr *errs.Error
}

func NewTokenManager() *TokenManager {
	return &TokenManager{
		key: []byte(config.TokenSigningKey),

		invalidTokenErr:  errs.NewError(errs.NotAuthorized, "invalid token"),
		notAuthorizedErr: errs.NewError(errs.NotAuthorized, "invalid token"),
	}
}

func (m *TokenManager) ValidateToken(token string) *errs.Error {
	_, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return consts.EmptyString, m.invalidTokenErr
		}
		return m.key, nil
	})
	if err != nil {
		return m.invalidTokenErr
	}
	return nil
}

func (m *TokenManager) ParseToken(token string) (string, *errs.Error) {
	t, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return consts.EmptyString, m.notAuthorizedErr
		}
		return m.key, nil
	})

	if err != nil || !t.Valid {
		return consts.EmptyString, m.notAuthorizedErr
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return consts.EmptyString, m.notAuthorizedErr
	}

	expTime, ok := claims["expTime"]
	if !ok {
		return consts.EmptyString, m.notAuthorizedErr
	}

	expTimeStr, ok := expTime.(string)
	if !ok {
		return consts.EmptyString, m.notAuthorizedErr
	}

	expTimeValue, err := time.Parse(time.RFC3339, expTimeStr)
	if err != nil {
		return consts.EmptyString, errs.NewServiceError(err)
	}

	if time.Now().After(expTimeValue) {
		return consts.EmptyString, m.notAuthorizedErr
	}

	user, ok := claims["userId"]
	if !ok {
		return consts.EmptyString, m.notAuthorizedErr
	}

	userId, ok := user.(string)
	if !ok {
		return consts.EmptyString, m.notAuthorizedErr
	}

	return userId, nil
}
