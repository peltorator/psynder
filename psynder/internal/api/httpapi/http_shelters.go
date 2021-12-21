package httpapi

import (
	"github.com/gorilla/mux"
	"github.com/peltorator/psynder/internal/api/httpapi/httperror"
	"github.com/peltorator/psynder/internal/api/httpapi/json"
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/auth"
	"github.com/peltorator/psynder/internal/domain/shelter"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/serviceimpl/authservice"
	"go.uber.org/zap"
	"net/http"
)

type httpApiShelters struct {
	authService    auth.Service
	shelterService shelter.Service
	jsonRW         json.ReadWriter
	eh             httperror.Handler
	logger         *zap.SugaredLogger
}

type ArgsShelters struct {
	DevMode        bool
	AuthService    auth.Service
	ShelterService shelter.Service
	Logger         *zap.SugaredLogger
}

func NewShelters(args ArgsShelters) *httpApiShelters {
	jsonRW := json.NewReadWriter()
	return &httpApiShelters{
		authService:    args.AuthService,
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

func (a *httpApiShelters) RouterShelters() http.Handler {
	r := mux.NewRouter()

	ar := r.NewRoute().Subrouter()
	ar.Use(a.authService.Authenticate)

	r.HandleFunc("/get-psyna-likes", a.eh.HandleErrors(a.psynaLikes)).Methods(http.MethodPost)
	ar.HandleFunc("/add-shelter-info", a.eh.HandleErrors(a.addShelterInfo)).Methods(http.MethodPost)
	ar.HandleFunc("/add-psyna", a.eh.HandleErrors(a.addPsyna)).Methods(http.MethodPost)
	ar.HandleFunc("/delete-psyna", a.eh.HandleErrors(a.deletePsyna)).Methods(http.MethodPost)
	withPaginationQueries(ar.HandleFunc("/browse-my-psynas", a.eh.HandleErrors(a.myPsynas))).Methods(http.MethodPost)

	return r
}

func (a *httpApiShelters) psynaLikes(w http.ResponseWriter, r *http.Request) error {
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

type addShelterInfoRequest struct {
	City    string `json:"city"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (a *httpApiShelters) addShelterInfo(w http.ResponseWriter, r *http.Request) error {
	var m addShelterInfoRequest
	err := a.jsonRW.ReadJson(r, &m)
	if err != nil {
		return err
	}
	acc := r.Context().Value(authservice.CtxUidKey).(domain.AccountId)

	err = a.shelterService.AddInfo(acc, domain.ShelterInfo{
		City:      m.City,
		Address:   m.Address,
		Phone:     m.Phone,
	})
	return err
}

type addPsynaRequest struct {
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Description string `json:"description"`
	PhotoLink   string `json:"photo_link"`
}

func (a *httpApiShelters) addPsyna(w http.ResponseWriter, r *http.Request) error {
	var m addPsynaRequest
	err := a.jsonRW.ReadJson(r, &m)
	if err != nil {
		return err
	}
	acc := r.Context().Value(authservice.CtxUidKey).(domain.AccountId)

	pid, err := a.shelterService.AddPsyna(acc, swipe.PsynaData{
		Name:        m.Name,
		Breed:       m.Breed,
		Description: m.Description,
		PhotoLink:   m.PhotoLink,
	})
	if err != nil {
		return err
	}
	return a.jsonRW.RespondWithJson(w, http.StatusOK, pid)
}

type deletePsynaRequest struct {
	Id uint64 `json:"id"`
}

func (a *httpApiShelters) deletePsyna(w http.ResponseWriter, r *http.Request) error {
	var m deletePsynaRequest
	err := a.jsonRW.ReadJson(r, &m)
	if err != nil {
		return err
	}
	acc := r.Context().Value(authservice.CtxUidKey).(domain.AccountId)

	return a.shelterService.DeletePsyna(acc, domain.PsynaId(m.Id))
}

func (a *httpApiShelters) myPsynas(w http.ResponseWriter, r *http.Request) error {
	acc := r.Context().Value(authservice.CtxUidKey).(domain.AccountId)

	pg, err := getPaginationInfo(r)
	if err != nil {
		return err
	}

	psynas, err := a.shelterService.GetMyPsynas(acc, pg)
	if err != nil {
		return err
	}

	return a.jsonRW.RespondWithJson(w, http.StatusOK, psynas)
}
