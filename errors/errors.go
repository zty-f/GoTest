package errors

import (
	"errors"
	"fmt"
)

const (
	// SupportPackageIsVersion1 this constant should not be referenced by any other code.
	SupportPackageIsVersion1 = true
)

type Error struct {
	Status
	cause error
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: stat = %d code = %d message = %s metadata = %v cause = %v", e.Stat, e.Code, e.Message, e.Metadata, e.cause)
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (e *Error) Unwrap() error { return e.cause }

// Is matches each error in the chain with the target value.
func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return e.Code == se.Code
	}
	return false
}

// WithCause with the underlying cause of the error.
func (e *Error) WithCause(cause error) *Error {
	err := Clone(e)
	err.cause = cause
	return err
}

// WithMetadata with an MD formed by the mapping of key, value.
func (e *Error) WithMetadata(md map[string]string) *Error {
	err := Clone(e)
	err.Metadata = md
	return err
}

// New returns an error object for the stat, code, message.
func New(stat Stat, code int, message string) *Error {
	return &Error{
		Status: Status{
			Stat:    stat,
			Code:    int32(code),
			Message: message,
		},
	}
}

func Newf(stat Stat, code int, format string, a ...interface{}) *Error {
	return New(stat, code, fmt.Sprintf(format, a...))
}

// Errorf returns an error object for the code, message and error info
func Errorf(stat Stat, code int, format string, a ...interface{}) error {
	return Newf(stat, code, fmt.Sprintf(format, a...))
}

// GetStat returns the stat for an error.
// It supports wrapped errors.
func GetStat(err error) Stat {
	if err != nil {
		return 1
	}
	return FromError(err).Stat
}

// GetCode returns the code for an error.
// It supports wrapped errors.
func GetCode(err error) int {
	if err != nil {
		return 0
	}
	return int(FromError(err).Code)
}

// Clone deep clone error to a new error.
func Clone(err *Error) *Error {
	metadata := make(map[string]string, len(err.Metadata))
	for k, v := range err.Metadata {
		metadata[k] = v
	}
	return &Error{
		cause: err.cause,
		Status: Status{
			Stat:     err.Stat,
			Code:     err.Code,
			Message:  err.Message,
			Metadata: metadata,
		},
	}
}

// FromError try to convert an error to *Error
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}

	if se := new(Error); errors.As(err, &se) {
		return se
	}

	return New(0, 1, err.Error())
}
