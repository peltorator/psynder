package httpapi

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"psynder/internal/usecases/account"
)

type Api struct {
	AccountUseCases account.AccountUseCasesInterface
}

func NewApi(a account.AccountUseCasesInterface) *Api {
	return &Api{
		AccountUseCases: a,
	}
}

func (a *Api) Router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/signup/giver", a.postSignupGiver).Methods(http.MethodPost)
	r.HandleFunc("/signup/taker", a.postSignupTaker).Methods(http.MethodPost)

	r.HandleFunc("/signin", a.postSignin).Methods(http.MethodPost)

	return r
}

type postSignupRequestModel struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// postSignup handles request for a new account creation.
func (a *Api) postSignupGiver(w http.ResponseWriter, r *http.Request) {
	var m postSignupRequestModel
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	acc, err := a.AccountUseCases.LoggerCreateAccount(a.AccountUseCases.CreateAccount)(m.Login, m.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Println(err)
		return
	}

	location := fmt.Sprintf("/accounts/%s", acc.Id)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

// postSignup handles request for a new account creation.
func (a *Api) postSignupTaker(w http.ResponseWriter, r *http.Request) {
	var m postSignupRequestModel
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	acc, err := a.AccountUseCases.LoggerCreateAccount(a.AccountUseCases.CreateAccount)(m.Login, m.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Println(err)
		return
	}

	location := fmt.Sprintf("/accounts/%s", acc.Id)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

// postSignin handles login request for existing user.
func (a *Api) postSignin(w http.ResponseWriter, r *http.Request) {
	var m postSignupRequestModel
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := a.AccountUseCases.LoggerLoginToAccount(a.AccountUseCases.LoginToAccount)(m.Login, m.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/jwt")
	if _, err := w.Write([]byte(token)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}