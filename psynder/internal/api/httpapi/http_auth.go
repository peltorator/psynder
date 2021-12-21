package httpapi

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/peltorator/psynder/internal/api/httpapi/httperror"
	"github.com/peltorator/psynder/internal/api/httpapi/json"
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/auth"
	"go.uber.org/zap"
	"net/http"
)

type httpApiAccounts struct {
	authService auth.Service
	jsonRW      json.ReadWriter
	eh          httperror.Handler
	logger      *zap.SugaredLogger
}

type ArgsAccounts struct {
	DevMode     bool
	AuthService auth.Service
	Logger      *zap.SugaredLogger
}

func NewAcccounts(args ArgsAccounts) *httpApiAccounts {
	jsonRW := json.NewReadWriter()
	return &httpApiAccounts{
		authService: args.AuthService,
		jsonRW:      jsonRW,
		eh: httperror.NewHandler(httperror.HandlerArgs{
			DevMode:        args.DevMode,
			JsonReadWriter: jsonRW,
			Logger:         args.Logger,
		}),
		logger: args.Logger,
	}
}

func (a *httpApiAccounts) RouterAccounts() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/signup", a.eh.HandleErrors(a.signup)).Methods(http.MethodPost)
	r.HandleFunc("/login", a.eh.HandleErrors(a.login)).Methods(http.MethodPost)

	return r
}

type signupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Kind     domain.AccountKind
}

func (a *httpApiAccounts) signup(w http.ResponseWriter, r *http.Request) error {
	var req signupRequest
	if err := a.jsonRW.ReadJson(r, &req); err != nil {
		return err
	}

	uid, err := a.authService.Signup(auth.SignupArgs{
		Credentials: auth.Credentials{
			Email:    req.Email,
			Password: req.Password,
		},
		Kind: req.Kind,
	})
	if err != nil {
		if errSignup, ok := err.(auth.SignupError); ok {
			statusCode, displayText := a.displaySignupError(errSignup)
			a.eh.RespondWithExpectedError(w, statusCode, displayText)
			return nil
		} else {
			return err
		}
	}

	location := fmt.Sprintf("/accounts/%s", uid.String())
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	return nil
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponseSuccess struct {
	Token string             `json:"token"`
	Kind  domain.AccountKind `json:"kind"`
}

func (a *httpApiAccounts) login(w http.ResponseWriter, r *http.Request) error {
	var req loginRequest
	if err := a.jsonRW.ReadJson(r, &req); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	tok, kind, err := a.authService.Login(auth.Credentials{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errLogin, ok := err.(auth.LoginError); ok {
			_ = errLogin // TODO!!!
			return err
		} else {
			return err
		}
	}

	w.Header().Set("Content-Type", "application/jwt")
	if err := a.jsonRW.WriteJson(w, loginResponseSuccess{Token: tok.String(), Kind: kind}); err != nil {
		return err
	}
	return nil
}
