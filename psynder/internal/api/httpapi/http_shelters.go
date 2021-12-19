package httpapi

import (
	"github.com/gorilla/mux"
	"github.com/peltorator/psynder/internal/api/httpapi/httperror"
	"github.com/peltorator/psynder/internal/api/httpapi/json"
	"github.com/peltorator/psynder/internal/domain/auth"
	"github.com/peltorator/psynder/internal/domain/shelter"
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
