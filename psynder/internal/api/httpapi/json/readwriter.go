package json

import (
	"github.com/peltorator/psynder/internal/errf"
	"net/http"
)

type ReadErrorKind int

const (
	ReadErrorUnknown = iota
	ReadErrorParse
)

type ReadError struct {
	Cause error
	Kind  ReadErrorKind
}

func (e ReadError) Error() string {
	return errf.WithKindAndCause("json read", int(e.Kind), e.Cause)
}

type WriteError struct {
	Cause error
	ValueToWrite interface{}
	ResponseStatusCode int
}

func (e WriteError) Error() string {
	return errf.WithCause("json write", e.Cause)
}

type ReadWriter interface {
	ReadJson(r *http.Request, val interface{}) error
	WriteJson(w http.ResponseWriter, val interface{}) error
	RespondWithJson(w http.ResponseWriter, statusCode int, val interface{}) error
}
