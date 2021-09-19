package usecases

import (
	"golang.org/x/crypto/bcrypt"
	"psynder/internal/domain/model"
	"psynder/internal/domain/repo"
	"psynder/internal/service/token"
	"unicode"
)

//var (
//	errInvalidLoginString    = errors.New("login string contains an invalid character")
//	errInvalidPasswordString = errors.New("password string contains an invalid character")
//	errPasswordTooShort      = errors.New("password is too short")
//	errTooLongString         = errors.New("too long string")
//	errNoCapitalLetters      = errors.New("password string does not contain capital letters")
//	errNoDigits              = errors.New("password string does not contain digits")
//)

const (
	minPasswordLength = 6
	maxPasswordLength = 50
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
	//GetAccountById(id string) (model.Account, error)
	//Authenticate(token string) (string, error)
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
	accId, err := u.AccountRepo.CreateAccount(repo.CreateAccountOptions{
		Email:        opts.Email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return 0, err
	}
	return accId, nil
}

//func (u *AccountUseCasesImpl) GetAccountById(id string) (Account, error) {
//	acc, err := u.AccountRepo.GetAccountById(id)
//	if err != nil {
//		return Account{}, err
//	}
//	return Account{Id: acc.Id}, err
//}

func (u *AccountUseCasesImpl) LoginToAccount(opts LoginToAccountOptions) (token.AccessToken, error) {
	id, err := u.AccountRepo.GetIdByEmail(opts.Email)
	if err != nil {
		return "", err
	}
	passwordHash, err := u.AccountRepo.GetPasswordHashById(id)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword(passwordHash, []byte(opts.Password)); err != nil {
		return "", err
	}
	tok, err := u.TokenIssuer.IssueToken(id)
	if err != nil {
		return "", err
	}
	return tok, err
}

//func (u *AccountUseCasesImpl) Authenticate(token string) (string, error) {
//	return u.TokenIssuer.AccountIdByToken(token)
//}

func validateEmail(email string) error {
	// TODO: just validate the email looool
	//chars := 0
	//for _, r := range login {
	//	if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
	//		return errInvalidLoginString
	//	}
	//	chars++
	//}
	//if chars < minLoginLength {
	//	return errTooShortString
	//}
	//if chars > maxLoginLength {
	//	return errTooLongString
	//}
	return nil
	//return fmt.Errorf("avtor servera lox")
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

//func (u *AccountUseCasesImpl) logger(method string, err error, start time.Time) {
//	status := "SUCCESS"
//	if err != nil {
//		status = err.Error()
//	}
//	fmt.Printf("method: %s; status-code: %s; call time: %v; duration: %v;\n",
//		method, status, start, time.Since(start))
//}

//func (u *AccountUseCasesImpl) LoggerCreateAccount(
//	createAccount func(login, password string) (Account, error)) func(login, password string) (Account, error) {
//
//	return func(login, password string) (Account, error) {
//		start := time.Now()
//		acc, err := createAccount(login, password)
//		u.logger("CreateAccount", err, start)
//		return acc, err
//	}
//}
//
//func (u *AccountUseCasesImpl) LoggerGetAccountById(
//	getAccountById func(id string) (Account, error)) func(id string) (Account, error) {
//
//	return func(id string) (Account, error) {
//		start := time.Now()
//		acc, err := getAccountById(id)
//		u.logger("GetAccountById", err, start)
//		return acc, err
//	}
//}
//
//func (u *AccountUseCasesImpl) LoggerLoginToAccount(
//	loginToAccount func(login, password string) (string, error)) func(login, password string) (string, error) {
//
//	return func(login, password string) (string, error) {
//		start := time.Now()
//		token, err := loginToAccount(login, password)
//		u.logger("LoginToAccount", err, start)
//		return token, err
//	}
//}
//
//func (u *AccountUseCasesImpl) LoggerAuthenticate(
//	authenticate func(token string) (string, error)) func(token string) (string, error) {
//
//	return func(token string) (string, error) {
//		start := time.Now()
//		token, err := authenticate(token)
//		u.logger("Authenticate", err, start)
//		return token, err
//	}
//}
