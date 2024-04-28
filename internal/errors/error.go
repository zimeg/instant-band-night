package errors

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// ToIBNError converts an error into a known IBNError
func ToIBNError(err error) IBNError {
	if ibn, ok := err.(IBNError); ok {
		return ibn
	}
	if err == huh.ErrUserAborted {
		return ErrPromptInterrupt
	}
	return IBNError{
		Message: err.Error(),
	}
}

// IBNError contains information about known program errors
type IBNError struct {
	Code    string
	Message string
}

// Error formats the error message for reader output
func (err IBNError) Error() string {
	return err.Message
}

// WithMessage updates the error message of an error
func (err IBNError) WithMessage(format string, a ...interface{}) IBNError {
	err.Message = fmt.Sprintf(format, a...)
	return err
}
