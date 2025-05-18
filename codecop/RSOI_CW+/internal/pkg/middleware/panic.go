package middleware

import (
	"encoding/json"
	"net/http"
)

func InternalServerError(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				answer := DefaultError
				jsn, _ := json.Marshal(answer)
				_, _ = w.Write(jsn)
			}
		}()

		handler.ServeHTTP(w, r)
	})
}
