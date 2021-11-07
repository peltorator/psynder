package httperror

import (
	"net/http"
)

type HandlerFuncWithError func(http.ResponseWriter, *http.Request) error

type Handler interface {
	HandleErrors(next HandlerFuncWithError) http.HandlerFunc
	RespondWithExpectedError(w http.ResponseWriter, statusCode int, displayText string)
	RespondWithUnexpectedError(w http.ResponseWriter, statusCode int, description string, debugInfo interface{})
}
