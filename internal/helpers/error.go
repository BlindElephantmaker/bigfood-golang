package helpers

import "errors"

const (
	badRequest = ErrorType(iota)
	unprocessableEntity
	//hideFromUser todo
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
	return newError(msg, badRequest)
}

func ErrorUnprocessableEntity(msg string) Error {
	return newError(msg, unprocessableEntity)
}

func newError(msg string, errorType ErrorType) Error {
	return Error{
		errorType:     errorType,
		originalError: errors.New(msg),
	}
}
