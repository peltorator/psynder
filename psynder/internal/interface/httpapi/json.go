package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONHandler interface {
	ReadJson(r *http.Request, val interface{}) error
	WriteJson(w http.ResponseWriter, val interface{}) error
}

type jsonParseError struct {
	Cause error
}

type jsonParseErrorResponse struct {
	Error string `json:"error"`
}

func newJSONParseError(cause error) *jsonParseError {
	return &jsonParseError{Cause: cause}
}

func (j *jsonParseError) ResponseData() interface{} {
	return jsonParseErrorResponse{
		// TODO: what if the schema is invalid???
		Error: fmt.Sprintf("Failed to parse request body: %s", j.Cause.Error()),
	}
}

func (j *jsonParseError) StatusCode() int {
	return http.StatusBadRequest
}

func (j *jsonParseError) Error() string {
	return j.Cause.Error()
}

type JSONHandlerImpl struct{}

func (j *JSONHandlerImpl) ReadJson(r *http.Request, val interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(val); err != nil {
		return newJSONParseError(err)
	}
	return nil
}

func (j *JSONHandlerImpl) WriteJson(w http.ResponseWriter, val interface{}) error {
	if err := json.NewEncoder(w).Encode(val); err != nil {
		return err
	}
	return nil
}
