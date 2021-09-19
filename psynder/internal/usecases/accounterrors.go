package usecases

import (
	"fmt"
	"net/http"
)

type accountCreationErrorKind int

const (
	errPasswordTooShort accountCreationErrorKind = iota
	errPasswordTooLong
	errPasswordInvalidChars
)

var errToDisplayedText = map[accountCreationErrorKind]string{
	errPasswordTooShort:     "Password is too short",
	errPasswordTooLong:      "Password is too long",
	errPasswordInvalidChars: "Password contains invalid characters",
}

type accountCreationError struct {
	Kind accountCreationErrorKind
}

func newAccountCreationError(kind accountCreationErrorKind) *accountCreationError {
	return &accountCreationError{
		Kind: kind,
	}
}

// AccountCreationErrorResponse TODO: this should not be exported if possible
type AccountCreationErrorResponse struct {
	Error string `json:"error"`
}

func (e *accountCreationError) ResponseData() interface{} {
	return AccountCreationErrorResponse{
		Error: errToDisplayedText[e.Kind],
	}
}

func (e *accountCreationError) StatusCode() int {
	return http.StatusBadRequest
}

func (e *accountCreationError) Error() string {
	// TODO: more debug info here?
	return fmt.Sprintf("failed to create an account: %s", errToDisplayedText[e.Kind])
}
