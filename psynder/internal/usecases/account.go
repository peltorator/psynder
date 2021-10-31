package usecases

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"net/smtp"
	"os"
	"github.com/peltorator/psynder/internal/domain/model"
	"github.com/peltorator/psynder/internal/domain/repo"
	"github.com/peltorator/psynder/internal/service/token"
	"unicode"
)

const (
	minPasswordLength = 6
	maxPasswordLength = 50
    smtpHost = "smtp.gmail.com"
    smtpPort = "587"
    psynderEmailAddress = "psynderapp@gmail.com"
)

type CreateAccountOptions struct {
	Email    string
	Password string
}

type LoginToAccountOptions struct {
	Email    string
	Password string
}

type AccountUseCases interface {
	CreateAccount(opts CreateAccountOptions) (model.AccountId, error)
	LoginToAccount(opts LoginToAccountOptions) (token.AccessToken, error)
	AuthenticateWithToken(token token.AccessToken) (model.AccountId, error)
}

type AccountUseCasesImpl struct {
	AccountRepo repo.AccountRepo
	TokenIssuer token.Issuer
}

func NewAccountUseCases(accountRepo repo.AccountRepo, tokenIssuer token.Issuer) *AccountUseCasesImpl {
	return &AccountUseCasesImpl{
		AccountRepo: accountRepo,
		TokenIssuer: tokenIssuer,
	}
}

func (u *AccountUseCasesImpl) CreateAccount(opts CreateAccountOptions) (model.AccountId, error) {
	if err := validateEmail(opts.Email); err != nil {
		return 0, err
	}
	if err := validatePassword(opts.Password); err != nil {
		return 0, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(opts.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
    //err = sendRegistrationEmail(opts.Email)
    //if err != nil {
    //    return 0, err
    //}
	accId, err := u.AccountRepo.StoreAccountToRepo(repo.CreateAccountOptions{
		Email:        opts.Email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return 0, err
	}
	return accId, nil
}

func (u *AccountUseCasesImpl) LoginToAccount(opts LoginToAccountOptions) (token.AccessToken, error) {
	id, err := u.AccountRepo.LoadIdByEmailFromRepo(opts.Email)
	if err != nil {
		return "", err
	}
	passwordHash, err := u.AccountRepo.LoadPasswordHashByIdFromRepo(id)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword(passwordHash, []byte(opts.Password)); err != nil {
		return "", newLoginError(errPasswordInvalid)
	}
	tok, err := u.TokenIssuer.IssueToken(id)
	if err != nil {
		return "", err
	}
	return tok, err
}

func (u *AccountUseCasesImpl) AuthenticateWithToken(token token.AccessToken) (model.AccountId, error) {
	return u.TokenIssuer.AccountIdByToken(token)
}

func validateEmail(email string) error {
    _, err := mail.ParseAddress(email)
    if err != nil {
        return newAccountCreationError(errEmailIncorrect)
    }
   return nil
}

func validatePassword(password string) error {
	chars := 0
	for _, r := range password {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			return newAccountCreationError(errPasswordInvalidChars)
		}
		chars++
	}
	if chars < minPasswordLength {
		return newAccountCreationError(errPasswordTooShort)
	}
	if chars > maxPasswordLength {
		return newAccountCreationError(errPasswordTooLong)
	}
	return nil
}

func sendRegistrationEmail(email string) error {
    auth := smtp.PlainAuth("", psynderEmailAddress, os.Getenv("emailpassword"), smtpHost)
    message := []byte("Thank you for registering on psynder")
    err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, psynderEmailAddress, []string{email}, message)
    return err
}
