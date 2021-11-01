package json

import (
	"encoding/json"
	"net/http"
)

type jsonReadWriter struct{}

func NewReadWriter() *jsonReadWriter {
	return &jsonReadWriter{}
}

func (j *jsonReadWriter) ReadJson(r *http.Request, val interface{}) error {
	// TODO: properly validate the schema here?
	if err := json.NewDecoder(r.Body).Decode(val); err != nil {
		return ReadError{Cause: err, Kind: ReadErrorParse}
	}
	return nil
}

func (j *jsonReadWriter) WriteJson(w http.ResponseWriter, val interface{}) error {
	if err := json.NewEncoder(w).Encode(val); err != nil {
		return WriteError{Cause: err, ValueToWrite: val}
	}
	return nil
}

func (j *jsonReadWriter) RespondWithJson(w http.ResponseWriter, statusCode int, val interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(val); err != nil {
		return WriteError{Cause: err, ValueToWrite: val, ResponseStatusCode: statusCode}
	}
	return nil
}

