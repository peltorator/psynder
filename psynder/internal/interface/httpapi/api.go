package httpapi

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"psynder/internal/usecases"
)

type Api struct {
	AccountUseCases usecases.AccountUseCases
	JSONHandler JSONHandler
}

func New(accountUseCases usecases.AccountUseCases, jsonHandler JSONHandler) *Api {
	return &Api{
		AccountUseCases: accountUseCases,
		JSONHandler: jsonHandler,
	}
}

func (a *Api) Router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/signup", handleErrorResponses(a.postSignup)).Methods(http.MethodPost)
	r.HandleFunc("/login", handleErrorResponses(a.postLogin)).Methods(http.MethodPost)

	return r
}

// postSignup handles request for a new account creation.
func (a *Api) postSignup(w http.ResponseWriter, r *http.Request) error {
	var m postSignupRequest
	if err := a.JSONHandler.ReadJson(r, &m); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	accId, err := a.AccountUseCases.CreateAccount(usecases.CreateAccountOptions{
		Email:    m.Email,
		Password: m.Password,
	})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	location := fmt.Sprintf("/accounts/%s", accId.String())
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	return nil
}

// postLogin handles login request for existing user.
func (a *Api) postLogin(w http.ResponseWriter, r *http.Request) error {
	var m postSignupRequest
	if err := a.JSONHandler.ReadJson(r, &m); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	tok, err := a.AccountUseCases.LoginToAccount(usecases.LoginToAccountOptions{
		Email:    m.Email,
		Password: m.Password,
	})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	w.Header().Set("Content-Type", "application/jwt")
	if err := a.JSONHandler.WriteJson(w, postLoginResponseSuccess{Token: tok.String()}); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	return nil
}
