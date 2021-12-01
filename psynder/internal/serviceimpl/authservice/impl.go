package authservice

import (
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/auth"
	"github.com/peltorator/psynder/internal/repo"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"unicode"
)

type AuthService struct {
	accRepo   repo.Accounts
	tokIssuer auth.TokenIssuer
}

func New(accountRepo repo.Accounts, tokenIssuer auth.TokenIssuer) *AuthService {
	return &AuthService{
		accRepo:   accountRepo,
		tokIssuer: tokenIssuer,
	}
}

func (a *AuthService) Signup(args auth.SignupArgs) (domain.AccountId, error) {
	const bcryptHashingCost = bcrypt.DefaultCost

	if args.Kind == domain.AccountKindUndefined {
		return 0, auth.SignupError{Kind: auth.SignupErrorAccountKindInvalid}
	}
	if err := validateEmail(args.Email); err != nil {
		return 0, err
	}
	if err := validatePassword(args.Password); err != nil {
		return 0, err
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(args.Password), bcryptHashingCost)
	if err != nil {
		return 0, auth.SignupError{
			Cause: err,
			Kind:  auth.SignupErrorUnknown,
		}
	}

	uid, err := a.accRepo.StoreNew(repo.AccountData{
		LoginCredentials: repo.LoginCredentials{
			Email:        args.Email,
			PasswordHash: passHash,
		},
		Kind: args.Kind,
	})
	if err != nil {
		errStore, ok := err.(repo.AccountStoreError)
		if !ok {
			return 0, err
		}

		var kind auth.SignupErrorKind
		switch errStore.Kind {
		case repo.AccountStoreErrorDuplicate:
			kind = auth.SignupErrorEmailTaken
		}

		return 0, auth.SignupError{
			Cause: errStore,
			Kind:  kind,
		}
	}

	return uid, nil
}

func (a *AuthService) Login(cred auth.Credentials) (auth.Token, domain.AccountKind, error) {
	acc, err := a.accRepo.LoadByEmail(cred.Email)
	if err != nil {
		errLoad, ok := err.(repo.AccountLoadError)
		if !ok {
			return "", domain.AccountKindUndefined, err
		}

		var kind auth.LoginErrorKind
		switch errLoad.Kind {
		case repo.AccountLoadErrorNoSuchEmail:
			kind = auth.LoginErrorNoMatchingAccount
		}

		return "", domain.AccountKindUndefined, auth.LoginError{
			Cause: errLoad,
			Kind:  kind,
		}
	}

	if err := bcrypt.CompareHashAndPassword(acc.PasswordHash, []byte(cred.Password)); err != nil {
		return "", domain.AccountKindUndefined, auth.LoginError{
			Cause: err,
			Kind:  auth.LoginErrorNoMatchingAccount,
		}
	}
	token, err := a.tokIssuer.IssueToken(acc.Id)
	return token, acc.Kind, err
}

func (a *AuthService) AuthByToken(tok auth.Token) (domain.AccountId, error) {
	uid, err := a.tokIssuer.AccountIdByToken(tok)
	if err != nil {
		return 0, auth.TokenError{
			Cause: err,
			Kind:  auth.TokenErrorInvalidToken,
		}
	}

	return uid, nil
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return auth.SignupError{
			Cause: err,
			Kind:  auth.SignupErrorEmailInvalid,
		}
	}
	return nil
}

func validatePassword(password string) error {
	const (
		minPasswordLength = 8
		maxPasswordLength = 40
	)

	chars := 0
	for _, r := range password {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			return auth.SignupError{Kind: auth.SignupErrorPasswordInvalidChars}
		}
		chars++
	}
	if chars < minPasswordLength {
		return auth.SignupError{Kind: auth.SignupErrorPasswordTooWeak}
	}
	if chars > maxPasswordLength {
		return auth.SignupError{Kind: auth.SignupErrorPasswordTooLong}
	}
	return nil
}
