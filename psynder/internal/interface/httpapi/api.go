package httpapi

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"psynder/internal/domain/model"
	"psynder/internal/domain/repo"
	"psynder/internal/usecases"
)

type Api struct {
	AccountUseCases usecases.AccountUseCases
	SwipeUseCases usecases.SwipeUseCases
	JSONHandler JSONHandler
}

func New(accountUseCases usecases.AccountUseCases, swipeUseCases usecases.SwipeUseCases, jsonHandler JSONHandler) *Api {
	return &Api{
		AccountUseCases: accountUseCases,
		SwipeUseCases: swipeUseCases,
		JSONHandler: jsonHandler,
	}
}

func (a *Api) Router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/signup", handleErrorResponses(a.postSignup)).Methods(http.MethodPost)
	r.HandleFunc("/login", handleErrorResponses(a.postLogin)).Methods(http.MethodPost)

	r.HandleFunc("/loadpsynas", handleErrorResponses(a.postLoadPsynas)).Methods(http.MethodPost)
	r.HandleFunc("/likepsyna", handleErrorResponses(a.postLikePsyna)).Methods(http.MethodPost)
	r.HandleFunc("/getfavoritepsynas", handleErrorResponses(a.postGetFavoritePsynas)).Methods(http.MethodPost)

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

func (a *Api) postLoadPsynas(w http.ResponseWriter, r *http.Request) error {
	var m postLoadPsynasRequest
	if err := a.JSONHandler.ReadJson(r, &m); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	psynas, err := a.SwipeUseCases.LoadPsynas(repo.LoadPsynasOptions{
		AccountId: model.AccountId(m.AccountId),
		Count: int (m.Count),
	})

	if err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	if err := a.JSONHandler.WriteJson(w, postLoadPsynasResponseSuccess{psynas}); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return nil
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func (a *Api) postLikePsyna(w http.ResponseWriter, r *http.Request) error {
	var m postLikePsynaRequest
	if err := a.JSONHandler.ReadJson(r, &m); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	err := a.SwipeUseCases.LikePsyna(repo.LikePsynaOptions{
		PsynaId: model.PsynaId(m.PsynaId),
		AccountId: model.AccountId(m.AccountId),
	})

	if err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (a *Api) postGetFavoritePsynas(w http.ResponseWriter, r *http.Request) error {
	var m postGetFavoritePsynasRequest
	if err := a.JSONHandler.ReadJson(r, &m); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	psynas, err := a.SwipeUseCases.GetFavoritePsynas(model.AccountId(m.AccountId))
	if err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	if err := a.JSONHandler.WriteJson(w, postGetFavoritePsynasResponseSuccess{psynas}); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return nil
	}
	w.WriteHeader(http.StatusOK)
	return nil
}
