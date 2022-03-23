package helpers

import "errors"

const (
	BadRequest   = ErrorType(iota)
	ShowToUser   // todo
	HideFromUser // todo
)

type ErrorType uint

type Error struct {
	errorType     ErrorType
	originalError error
}

func (e Error) Error() string {
	return e.originalError.Error()
}

func ErrorBadRequest(msg string) Error {
	return newError(msg, BadRequest)
}

func newError(msg string, errorType ErrorType) Error {
	return Error{
		errorType:     errorType,
		originalError: errors.New(msg),
	}
}
