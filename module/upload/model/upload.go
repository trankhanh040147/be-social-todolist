package model

import (
	"errors"
	"social-todo-list/common"
)

const (
	EntityName = "Upload"
)

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string { return "uploads" }

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge",
	)
)

func ErrFileNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"File is not image",
		"ErrFileNotImage",
	)
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot save file",
		"ErrCannotSaveFile",
	)
}
