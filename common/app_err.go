package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// AppError implements the error interface.
// AppError(AppError(...(error))) will return the root error.
// Example:
// err := errors.New("cannot connect to database server")
// appErr := NewErrorResponse(err, "cannot connect to database server", "cannot connect to database server", "cannot_connect_to_database_server")
// appErr.RootError() // 'cannot connect to database server

func (e *AppError) RootError() error {
	// If the root error is an AppError, return its root error instead.
	// This is useful when we want to wrap an error with a more specific error message.
	// For example, we want to wrap a database error with a message 'cannot connect to database
	// server' instead of 'pq: cannot connect to database server
	// Cast the root error to AppError and return its root error.
	// Otherwise, return the root error itself.

	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewFullErrorResponse(StatusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: StatusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"Something went wrong in the server", err.Error(), "ErrInternal")
}

func ErrValidation(err error) *AppError {
	return NewErrorResponse(err, "Validation error", err.Error(), "ErrValidation")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrEntityDeleted(entity string) *AppError {
	return NewCustomError(
		nil,
		fmt.Sprintf("%s has been deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

func ErrEntityExisted(entity string) *AppError {
	return NewCustomError(
		nil,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Err%sExisted", entity),
	)
}

func ErrEntityNotFound(entity string) *AppError {
	return NewCustomError(
		nil,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("You don't have permission to do this action"),
		fmt.Sprintf("ErrNoPermission"),
	)
}

var RecordNotFound = errors.New("record not found")
