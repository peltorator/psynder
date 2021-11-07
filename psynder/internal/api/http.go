package api

import (
	"net/http"
)

type Http interface {
	Router() http.Handler
}
