package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse interface {
	error
	ResponseData() interface{}
	StatusCode() int
}

func handleErrorResponses(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			errResponse, ok := err.(ErrorResponse)
			if ok {
				w.WriteHeader(errResponse.StatusCode())
				if err := json.NewEncoder(w).Encode(errResponse.ResponseData()); err != nil {
					// TODO: do not panic!
					panic(err)
				}
				return
			}
			// TODO: do not panic!
			panic(fmt.Errorf("unhandled error from handler: %w", err))
		}
	}
}