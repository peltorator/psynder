package auth

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/errf"
)

type SignupErrorKind int

const (
	SignupErrorUnknown SignupErrorKind = iota
	SignupErrorEmailTaken
	SignupErrorEmailInvalid
	SignupErrorPasswordInvalidChars
	SignupErrorPasswordTooLong
	SignupErrorPasswordTooWeak
	SignupErrorAccountKindInvalid
)

type SignupError struct {
	Cause error
	Kind  SignupErrorKind
}

func (e SignupError) Error() string {
	return errf.WithKindAndCause("signup", int(e.Kind), e.Cause)
}

type LoginErrorKind int

const (
	LoginErrorUnknown LoginErrorKind = iota
	LoginErrorNoMatchingAccount
)

type LoginError struct {
	Cause error
	Kind  LoginErrorKind
}

func (e LoginError) Error() string {
	return errf.WithKindAndCause("login", int(e.Kind), e.Cause)
}

type TokenErrorKind int

const (
	TokenErrorUnknown TokenErrorKind = iota
	TokenErrorInvalidToken
)

type TokenError struct {
	Cause error
	Kind  TokenErrorKind
}

func (e TokenError) Error() string {
	return errf.WithKindAndCause("token", int(e.Kind), e.Cause)
}

type SignupArgs struct {
	Credentials
	Kind domain.AccountKind
}

type Service interface {
	Signup(args SignupArgs) (domain.AccountId, error)
	Login(cred Credentials) (Token, error)
	AuthByToken(tok Token) (domain.AccountId, error)
}

type TokenIssuer interface {
	IssueToken(uid domain.AccountId) (Token, error)
	AccountIdByToken(tok Token) (domain.AccountId, error)
}
