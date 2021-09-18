package httpapi

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func (a *Api) authenticate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearHeader := r.Header.Get("Authorization")
		strArr := strings.Split(bearHeader, " ")
		if len(strArr) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token := strArr[1]
		id, err := a.AccountUseCases.LoggerAuthenticate(a.AccountUseCases.Authenticate)(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "account_id", id)
		handler(w, r.WithContext(ctx))
	}
}

type responseWriterObserver struct {
	http.ResponseWriter
	status int
	wroteHeader bool
}

func (o *responseWriterObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.status = code
}

func (o *responseWriterObserver) StatusCode() int {
	if !o.wroteHeader {
		return http.StatusOK
	}
	return o.status
}

func (a *Api) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		o := &responseWriterObserver{ResponseWriter: w}
		next.ServeHTTP(o, r)
		fmt.Printf("method: %s; status-code: %d; url: %s; remote-addr: %s; request call time: %v; duration: %v;\n",
			r.Method, o.StatusCode(), r.URL.String(), r.RemoteAddr, start, time.Since(start))
	})
}