package httpapi

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Api struct {
	
}

func New() *Api {
	return &Api{}
}

func (a *Api) Router() http.Handler {
	r := mux.NewRouter()

	return r
}