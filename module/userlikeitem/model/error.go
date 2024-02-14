package model

import "social-todo-list/common"

func ErrCannotLikeItem(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot like this item",
		"ErrCannotLikeItem",
	)
}

func ErrCannotUnlikeItem(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"Cannot unlike this item",
		"ErrCannotUnlikeItem",
	)
}

func ErrDidNotLikeItem(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"You have not liked this item",
		"ErrDidNotLikeItem",
	)
}
