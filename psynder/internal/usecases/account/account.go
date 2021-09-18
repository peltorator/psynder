package account

import (
"fmt"
"psynder/internal/domain/account"
"psynder/internal/service/token"
"time"

"errors"
"golang.org/x/crypto/bcrypt"
"unicode"
)

var (
	ErrInvalidLoginString    = errors.New("login string contains invalid character")
	ErrInvalidPasswordString = errors.New("password string contains invalid character")
	ErrTooShortString        = errors.New("too short string")
	ErrTooLongString         = errors.New("too long string")
	ErrNoCapitalLetters      = errors.New("password string does not contain capital letters")
	ErrNoDigits              = errors.New("password string does not contain digits")
)

const (
	minLoginLength    = 4
	maxLoginLength    = 50
	minPasswordLength = 6
	maxPasswordLength = 50
)

type Account struct {
	Id string
}

type AccountUseCasesInterface interface {
	CreateAccount(login, password string) (Account, error)
	GetAccountById(id string) (Account, error)
	LoginToAccount(login, password string) (string, error)
	Authenticate(token string) (string, error)

	//Logging
	LoggerCreateAccount(
		createAccount func(login, password string) (Account, error)) func(login, password string) (Account, error)
	LoggerGetAccountById(
		getAccountById func(id string) (Account, error)) func(id string) (Account, error)
	LoggerLoginToAccount(
		loginToAccount func(login, password string) (string, error)) func(login, password string) (string, error)
	LoggerAuthenticate(
		authenticate func(token string) (string, error)) func(token string) (string, error)
}

type AccountUseCases struct {
	AccountStorage account.Interface
	Auth           token.Interface
}

func (a *AccountUseCases) CreateAccount(login, password string) (Account, error) {
	if err := validateLogin(login); err != nil {
		return Account{}, err
	}
	if err := validatePassword(password); err != nil {
		return Account{}, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return Account{}, err
	}
	acc, err := a.AccountStorage.CreateAccount(account.Credentials{
		Login:    login,
		Password: string(hashedPassword),
	})
	if err != nil {
		return Account{}, err
	}
	return Account{Id: acc.Id}, nil
}

func (a *AccountUseCases) GetAccountById(id string) (Account, error) {
	acc, err := a.AccountStorage.GetAccountById(id)
	if err != nil {
		return Account{}, err
	}
	return Account{Id: acc.Id}, err
}

func (a *AccountUseCases) LoginToAccount(login, password string) (string, error) {
	if err := validateLogin(login); err != nil {
		return "", err
	}
	if err := validatePassword(password); err != nil {
		return "", err
	}
	acc, err := a.AccountStorage.GetAccountByLogin(login)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(acc.Credentials.Password), []byte(password)); err != nil {
		return "", err
	}
	token, err := a.Auth.IssueToken(acc.Id)
	if err != nil {
		return "", err
	}
	return token, err
}

func (a *AccountUseCases) Authenticate(token string) (string, error) {
	return a.Auth.UserIdByToken(token)
}

func validateLogin(login string) error {
	chars := 0
	for _, r := range login {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return ErrInvalidLoginString
		}
		chars++
	}
	if chars < minLoginLength {
		return ErrTooShortString
	}
	if chars > maxLoginLength {
		return ErrTooLongString
	}
	return nil
}

func validatePassword(password string) error {
	chars := 0
	digits := 0
	capitalLetters := 0
	for _, r := range password {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			return ErrInvalidPasswordString
		}

		if unicode.IsDigit(r) {
			digits += 1
		}
		if unicode.IsUpper(r) {
			capitalLetters += 1
		}

		chars++
	}
	if chars < minPasswordLength {
		return ErrTooShortString
	}
	if chars > maxPasswordLength {
		return ErrTooLongString
	}
	if digits == 0 {
		return ErrNoDigits
	}
	if capitalLetters == 0 {
		return ErrNoCapitalLetters
	}
	return nil
}

func (a *AccountUseCases) logger(method string, err error, start time.Time) {
	status := "SUCCESS"
	if err != nil {
		status = err.Error()
	}
	fmt.Printf("method: %s; status-code: %s; call time: %v; duration: %v;\n",
		method, status, start, time.Since(start))
}

func (a *AccountUseCases) LoggerCreateAccount(
	createAccount func(login, password string) (Account, error)) func(login, password string) (Account, error) {

	return func(login, password string) (Account, error) {
		start := time.Now()
		acc, err := createAccount(login, password)
		a.logger("CreateAccount", err, start)
		return acc, err
	}
}

func (a *AccountUseCases) LoggerGetAccountById(
	getAccountById func(id string) (Account, error)) func(id string) (Account, error) {

	return func(id string) (Account, error) {
		start := time.Now()
		acc, err := getAccountById(id)
		a.logger("GetAccountById", err, start)
		return acc, err
	}
}

func (a *AccountUseCases) LoggerLoginToAccount(
	loginToAccount func(login, password string) (string, error)) func(login, password string) (string, error) {

	return func(login, password string) (string, error) {
		start := time.Now()
		token, err := loginToAccount(login, password)
		a.logger("LoginToAccount", err, start)
		return token, err
	}
}

func (a *AccountUseCases) LoggerAuthenticate(
	authenticate func(token string) (string, error)) func(token string) (string, error) {

	return func(token string) (string, error) {
		start := time.Now()
		token, err := authenticate(token)
		a.logger("Authenticate", err, start)
		return token, err
	}
}

