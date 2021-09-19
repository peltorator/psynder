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
	errEmailTaken
)

var accountCreationToDisplayedText = map[accountCreationErrorKind]string{
	errPasswordTooShort:     "Password is too short",
	errPasswordTooLong:      "Password is too long",
	errPasswordInvalidChars: "Password contains invalid characters",
	errEmailTaken:           "A user with that email address already exists",
}

type accountCreationError struct {
	Kind accountCreationErrorKind
}

func newAccountCreationError(kind accountCreationErrorKind) *accountCreationError {
	return &accountCreationError{
		Kind: kind,
	}
}

// TODO: молодой человек у вас абстракция протекает...
type accountCreationErrorResponse struct {
	Error string `json:"error"`
}

func (e *accountCreationError) ResponseData() interface{} {
	return accountCreationErrorResponse{
		Error: accountCreationToDisplayedText[e.Kind],
	}
}

func (e *accountCreationError) StatusCode() int {
	return http.StatusBadRequest
}

func (e *accountCreationError) Error() string {
	// TODO: more debug info here?
	return fmt.Sprintf("failed to create an account: %s", accountCreationToDisplayedText[e.Kind])
}

type loginErrorKind int

const (
	errPasswordInvalid loginErrorKind = iota
)

var loginErrorToDisplayedText = map[loginErrorKind]string{
	errPasswordInvalid: "Incorrect password",
}

type loginError struct {
	Kind loginErrorKind
}

func newLoginError(kind loginErrorKind) *loginError {
	return &loginError{Kind: kind}
}

type LoginErrorResponse struct {
	Error string `json:"error"`
}

func (e *loginError) ResponseData() interface{} {
	return LoginErrorResponse{
		Error: loginErrorToDisplayedText[e.Kind],
	}
}

func (e *loginError) StatusCode() int {
	return http.StatusForbidden
}

func (e *loginError) Error() string {
	return fmt.Sprintf("login failed: %s", loginErrorToDisplayedText[e.Kind])
}
