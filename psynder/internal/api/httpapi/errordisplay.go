package httpapi

import (
	"github.com/peltorator/psynder/internal/domain/auth"
	"net/http"
)

func (a *httpApiAccounts) displaySignupError(err auth.SignupError) (int, string) {
	switch err.Kind {
	case auth.SignupErrorEmailTaken:
		return http.StatusConflict, "There already exists an account with this email"
	case auth.SignupErrorEmailInvalid:
		return http.StatusUnprocessableEntity, "Please enter a valid email address"
	case auth.SignupErrorPasswordInvalidChars:
		return http.StatusUnprocessableEntity, "Only alphanumeric characters and spaces are allowed in passwords"
	case auth.SignupErrorPasswordTooLong:
		return http.StatusUnprocessableEntity, "Password is too long"
	case auth.SignupErrorPasswordTooWeak:
		return http.StatusUnprocessableEntity, "Password is too weak"
	case auth.SignupErrorAccountKindInvalid:
		return http.StatusUnprocessableEntity, "Account kind should be one of: person, shelter"
	default:
		a.logger.DPanicf("Unknown signup error: %v", err)
		return http.StatusInternalServerError, "Unknown signup error"
	}
}

func (a *httpApiAccounts) displayLoginError(err auth.LoginError) (int, string) {
	switch err.Kind {
	case auth.LoginErrorNoMatchingAccount:
		return http.StatusForbidden, "There is no account matching this email and password"
	default:
		a.logger.DPanicf("Unknown login error: %v", err)
		return http.StatusInternalServerError, "Unknown login error"
	}
}
