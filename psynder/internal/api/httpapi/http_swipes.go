package httpapi

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/peltorator/psynder/internal/api/httpapi/httperror"
	"github.com/peltorator/psynder/internal/api/httpapi/json"
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/auth"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/serviceimpl/authservice"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type httpApiSwipes struct {
	authService  auth.Service
	swipeService swipe.Service
	jsonRW       json.ReadWriter
	eh           httperror.Handler
	logger       *zap.SugaredLogger
}

type Args struct {
	DevMode      bool
	AuthService  auth.Service
	SwipeService swipe.Service
	Logger       *zap.SugaredLogger
}

func New(args Args) *httpApiSwipes {
	jsonRW := json.NewReadWriter()
	return &httpApiSwipes{
		authService:  args.AuthService,
		swipeService: args.SwipeService,
		jsonRW:       jsonRW,
		eh: httperror.NewHandler(httperror.HandlerArgs{
			DevMode:        args.DevMode,
			JsonReadWriter: jsonRW,
			Logger:         args.Logger,
		}),
		logger: args.Logger,
	}
}

func (a *httpApiSwipes) Router() http.Handler {
	r := mux.NewRouter()

	ar := r.NewRoute().Subrouter()
	ar.Use(a.authService.Authenticate)

	withPaginationQueries(ar.HandleFunc("/browse-psynas", a.eh.HandleErrors(a.browsePsynas))).Methods(http.MethodGet)
	// TODO: handle no-params case correctly (error-handling)

	ar.HandleFunc("/like-psyna", a.eh.HandleErrors(a.likePsyna)).Methods(http.MethodPost)
	withPaginationQueries(ar.HandleFunc("/liked-psynas", a.eh.HandleErrors(a.getLikedPsynas))).Methods(http.MethodGet)

	r.HandleFunc("/psyna-info", a.eh.HandleErrors(a.psynaInfo)).Methods(http.MethodPost)
	r.HandleFunc("/get-all-info", a.eh.HandleErrors(a.allInfo)).Methods(http.MethodPost)

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

type browsePsynasRequest struct {
	Breed       *string           `json:"breed,omitempty"`
	ShelterCity *string           `json:"shelter_city,omitempty"`
	Shelter     *domain.AccountId `json:"shelter,omitempty"`
}

func (a *httpApiSwipes) browsePsynas(w http.ResponseWriter, r *http.Request) error {
	var m browsePsynasRequest
	//err := a.jsonRW.ReadJson(r, &m)
	//if err != nil {
	//	return err
	//}

	uid := r.Context().Value(authservice.CtxUidKey).(domain.AccountId)

	pg, err := getPaginationInfo(r)
	if err != nil {
		return err // TODO: handle this!
	}

	psynas, err := a.swipeService.BrowsePsynas(uid, pg, domain.PsynaFilter{
		Breed:       m.Breed,
		Shelter:     m.Shelter,
		ShelterCity: m.ShelterCity,
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

func (a *httpApiSwipes) likePsyna(w http.ResponseWriter, r *http.Request) error {
	acc := r.Context().Value(authservice.CtxUidKey).(domain.AccountId)

	var m likePsynaRequest
	if err := a.jsonRW.ReadJson(r, &m); err != nil {
		return err
	}

	if err := a.swipeService.RatePsyna(acc, m.PsynaId, swipe.DecisionLike); err != nil {
		return err
	}

	return nil
}

func (a *httpApiSwipes) getLikedPsynas(w http.ResponseWriter, r *http.Request) error {
	acc := r.Context().Value(authservice.CtxUidKey).(domain.AccountId)
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

func (a *httpApiSwipes) psynaInfo(w http.ResponseWriter, r *http.Request) error {
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

func (a *httpApiSwipes) allInfo(w http.ResponseWriter, r *http.Request) error {

	info, err := a.swipeService.GetAllInfo()

	if err != nil {
		return err
	}

	return a.jsonRW.RespondWithJson(w, http.StatusOK, info)
}
