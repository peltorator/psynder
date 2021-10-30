package httpapi

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"psynder/internal/domain/model"
	"psynder/internal/domain/repo"
	"psynder/internal/service/token"
	"psynder/internal/usecases"
	"regexp"
)

const (
	CTX_ACCOUNT_ID_KEY = iota
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

	r.HandleFunc("/signup", handleErrorResponses(a.signup)).Methods(http.MethodPost)
	r.HandleFunc("/login", handleErrorResponses(a.login)).Methods(http.MethodPost)

	ar := r.NewRoute().Subrouter()
	ar.Use(a.authenticate)

	ar.HandleFunc("/loadpsynas", handleErrorResponses(a.loadPsynas)).Methods(http.MethodPost)
	ar.HandleFunc("/likepsyna", handleErrorResponses(a.likePsyna)).Methods(http.MethodPost)
	ar.HandleFunc("/getfavoritepsynas", handleErrorResponses(a.getFavoritePsynas)).Methods(http.MethodGet)

	return r
}

// signup handles request for a new account creation.
func (a *Api) signup(w http.ResponseWriter, r *http.Request) error {
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

// login handles login request for existing user.
func (a *Api) login(w http.ResponseWriter, r *http.Request) error {
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

func (a *Api) loadPsynas(w http.ResponseWriter, r *http.Request) error {
	var m postLoadPsynasRequest
	if err := a.JSONHandler.ReadJson(r, &m); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	psynas, err := a.SwipeUseCases.LoadPsynas(repo.LoadPsynasOptions{
		AccountId: r.Context().Value(CTX_ACCOUNT_ID_KEY).(model.AccountId),
		Limit: int (m.Limit),
		Offset: int (m.Offset),
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

func (a *Api) likePsyna(w http.ResponseWriter, r *http.Request) error {
	var m postLikePsynaRequest
	if err := a.JSONHandler.ReadJson(r, &m); err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	err := a.SwipeUseCases.LikePsyna(repo.LikePsynaOptions{
		PsynaId: model.PsynaId(m.PsynaId),
		AccountId: r.Context().Value(CTX_ACCOUNT_ID_KEY).(model.AccountId),
	})

	if err != nil {
		fmt.Println("Error: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func (a *Api) getFavoritePsynas(w http.ResponseWriter, r *http.Request) error {
	psynas, err := a.SwipeUseCases.GetFavoritePsynas(r.Context().Value(CTX_ACCOUNT_ID_KEY).(model.AccountId))
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

var bearerTokenRegexp = regexp.MustCompile("Bearer (.*)")

func (a *Api) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		submatches := bearerTokenRegexp.FindStringSubmatch(authHeader)
		if len(submatches) != 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		tok := token.AccessToken(submatches[1])
		accId, err := a.AccountUseCases.AuthenticateWithToken(tok)
		if err != nil {
			//TODO: a che delat
			w.WriteHeader(http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), CTX_ACCOUNT_ID_KEY, accId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}