package tokenprovider

import (
	"errors"

	"social-todo-list/common"
)

var (
	ErrTokenNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrTokenNotFound",
	)
	ErrInvalidToken = common.NewCustomError(
		errors.New("token is invalid"),
		"token is invalid",
		"ErrInvalidToken",
	)
	ErrExpiredToken = common.NewCustomError(
		errors.New("token has expired"),
		"token has expired",
		"ErrExpiredToken",
	)
	ErrTokenEncoding = common.NewCustomError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrTokenEncoding",
	)
)

type TokenProvider interface {
	Generate(data TokenPayload, expiry int) (Token, error)
	Validate(token string) (TokenPayload, error)
	SecretKey() string
}

type TokenPayload interface {
	UserId() int
	Role() string
}

type Token interface {
	GetToken() string
}
