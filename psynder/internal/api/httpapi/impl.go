package httpapi

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/peltorator/psynder/internal/api/httpapi/httperror"
	"github.com/peltorator/psynder/internal/api/httpapi/json"
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/auth"
	"github.com/peltorator/psynder/internal/domain/shelter"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"go.uber.org/zap"
	"net/http"
	"regexp"
	"strconv"
)

const (
	ctxUidKey = iota + 10 // TODO: does this clash with mux?
)

type httpApi struct {
	authService    auth.Service
	swipeService   swipe.Service
	shelterService shelter.Service
	jsonRW         json.ReadWriter
	eh             httperror.Handler
	logger         *zap.SugaredLogger
}

type Args struct {
	DevMode        bool
	AuthService    auth.Service
	SwipeService   swipe.Service
	ShelterService shelter.Service
	Logger         *zap.SugaredLogger
}

func New(args Args) *httpApi {
	jsonRW := json.NewReadWriter()
	return &httpApi{
		authService:    args.AuthService,
		swipeService:   args.SwipeService,
		shelterService: args.ShelterService,
		jsonRW:         jsonRW,
		eh: httperror.NewHandler(httperror.HandlerArgs{
			DevMode:        args.DevMode,
			JsonReadWriter: jsonRW,
			Logger:         args.Logger,
		}),
		logger: args.Logger,
	}
}

func (a *httpApi) Router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/signup", a.eh.HandleErrors(a.signup)).Methods(http.MethodPost)
	r.HandleFunc("/login", a.eh.HandleErrors(a.login)).Methods(http.MethodPost)

	ar := r.NewRoute().Subrouter()
	ar.Use(a.authenticate)

	withPaginationQueries(ar.HandleFunc("/browse-psynas", a.eh.HandleErrors(a.browsePsynas))).Methods(http.MethodGet)
	// TODO: handle no-params case correctly (error-handling)

	ar.HandleFunc("/like-psyna", a.eh.HandleErrors(a.likePsyna)).Methods(http.MethodPost)
	withPaginationQueries(ar.HandleFunc("/liked-psynas", a.eh.HandleErrors(a.getLikedPsynas))).Methods(http.MethodGet)

	// TODO(antoha): add shelter http api

	r.HandleFunc("/psyna-info", a.eh.HandleErrors(a.psynaInfo)).Methods(http.MethodPost)
	r.HandleFunc("/get-psyna-likes", a.eh.HandleErrors(a.psynaLikes)).Methods(http.MethodPost)
	r.HandleFunc("/get-all-info", a.eh.HandleErrors(a.allInfo)).Methods(http.MethodPost)

	//ar.HandleFunc("/likepsyna", handleErrors(a.likePsyna)).Methods(http.MethodPost)
	//ar.HandleFunc("/getfavoritepsynas", handleErrors(a.getFavoritePsynas)).Methods(http.MethodGet)

	return r
}

func withPaginationQueries(r *mux.Route) *mux.Route {
	return r.Queries(
		"limit", "{limit:[0-9]+}",
		"offset", "{offset:[0-9]+}",
	)
}

// TODO: rewrite this function
func getPaginationInfo(r *http.Request) (pagination.Info, error) {
	vars := mux.Vars(r)

	var (
		limitStr, offsetStr string
		limit, offset       int
	)

	limitStr, ok := vars["limit"]
	if !ok {
		goto err
	}
	if limit64, err := strconv.ParseInt(limitStr, 10, 64); err != nil {
		goto err
	} else {
		limit = int(limit64)
	}

	offsetStr, ok = vars["offset"]
	if !ok {
		goto err
	}
	if offset64, err := strconv.ParseInt(offsetStr, 10, 64); err != nil {
		goto err
	} else {
		offset = int(offset64)
	}

	return pagination.Info{
		Limit:  limit,
		Offset: offset,
	}, nil
err:
	return pagination.Info{}, fmt.Errorf("limit and offset expected") // TODO: better error here

}

type signupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Kind     domain.AccountKind
}

func (a *httpApi) signup(w http.ResponseWriter, r *http.Request) error {
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

func (a *httpApi) login(w http.ResponseWriter, r *http.Request) error {
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

type browsePsynasRequest struct {
	SpecificBreed       *string           `json:"specific_breed,omitempty"`
	SpecificShelterCity *string           `json:"specific_shelter_city,omitempty"`
	SpecificShelter     *domain.AccountId `json:"specific_shelter,omitempty"`
}

func (a *httpApi) browsePsynas(w http.ResponseWriter, r *http.Request) error {
	var m browsePsynasRequest
	err := a.jsonRW.ReadJson(r, &m)
	if err != nil {
		return err
	}
	uid := r.Context().Value(ctxUidKey).(domain.AccountId)

	pg, err := getPaginationInfo(r)
	if err != nil {
		return err // TODO: handle this!
	}

	psynas, err := a.swipeService.BrowsePsynas(uid, pg, domain.PsynaFilter{
		SpecificBreed:       m.SpecificBreed,
		SpecificShelter:     m.SpecificShelter,
		SpecificShelterCity: m.SpecificShelterCity,
	})
	if err != nil {
		return err // TODO: handle this somehow?
	}

	if err := a.jsonRW.RespondWithJson(w, http.StatusOK, psynas); err != nil {
		return err // TODO: also set appropriate header
	}

	return nil
}

type likePsynaRequest struct {
	PsynaId domain.PsynaId `json:"psynaId"`
}

type psynaInfoRequest struct {
	PsynaId domain.PsynaId `json:"psynaId"`
}

type psynaLikesRequest struct {
	PsynaId domain.PsynaId `json:"psynaId"`
}

func (a *httpApi) likePsyna(w http.ResponseWriter, r *http.Request) error {
	acc := r.Context().Value(ctxUidKey).(domain.AccountId)

	var m likePsynaRequest
	if err := a.jsonRW.ReadJson(r, &m); err != nil {
		return err
	}

	if err := a.swipeService.RatePsyna(acc, m.PsynaId, swipe.DecisionLike); err != nil {
		return err
	}

	return nil
}

func (a *httpApi) getLikedPsynas(w http.ResponseWriter, r *http.Request) error {
	acc := r.Context().Value(ctxUidKey).(domain.AccountId)
	pg, err := getPaginationInfo(r)
	if err != nil {
		return err
	}

	likedPsynas, err := a.swipeService.GetLikedPsynas(acc, pg)
	if err != nil {
		return err
	}

	return a.jsonRW.RespondWithJson(w, http.StatusOK, likedPsynas)
}

func (a *httpApi) psynaInfo(w http.ResponseWriter, r *http.Request) error {
	var m psynaInfoRequest
	err := a.jsonRW.ReadJson(r, &m)
	if err != nil {
		return err
	}

	shelterInformation, err := a.swipeService.GetPsynaInfo(m.PsynaId)

	if err != nil {
		return err
	}

	return a.jsonRW.RespondWithJson(w, http.StatusOK, shelterInformation)
}

func (a *httpApi) psynaLikes(w http.ResponseWriter, r *http.Request) error {
	var m psynaLikesRequest
	err := a.jsonRW.ReadJson(r, &m)
	if err != nil {
		return err
	}

	likes, err := a.shelterService.GetPsynaLikes(m.PsynaId)

	if err != nil {
		return err
	}

	return a.jsonRW.RespondWithJson(w, http.StatusOK, likes)
}

func (a *httpApi) allInfo(w http.ResponseWriter, r *http.Request) error {

	info, err := a.swipeService.GetAllInfo()

	if err != nil {
		return err
	}

	return a.jsonRW.RespondWithJson(w, http.StatusOK, info)
}

var bearerTokenRegexp = regexp.MustCompile("Bearer (.*)")

func (a *httpApi) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		submatches := bearerTokenRegexp.FindStringSubmatch(authHeader)
		if len(submatches) != 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		tok := submatches[1]
		uid, err := a.authService.AuthByToken(auth.NewTokenFromString(tok))
		if err != nil {
			if errToken, ok := err.(auth.TokenError); ok && errToken.Kind == auth.TokenErrorInvalidToken {
				w.Header().Set("WWW-Authenticate", `Bearer error="invalid_token"`)
				w.WriteHeader(http.StatusUnauthorized)
			} else {
				a.logger.DPanicf("Unknown auth by token error: %v", err)
			}
			return
		}
		ctx := context.WithValue(r.Context(), ctxUidKey, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
