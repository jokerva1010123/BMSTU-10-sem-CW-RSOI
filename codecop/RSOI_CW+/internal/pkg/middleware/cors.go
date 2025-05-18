package middleware

import (
	"net/http"
)

const defaultOrigin = "http://3.67.182.34:8887"//"http://3.67.182.34:8887"

func Cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", defaultOrigin)
		w.Header().Set("Access-Control-Allow-Headers",
			"Authorization, Access-Control-Allow-Origin, " +
			"Access-Control-Allow-Headers, Origin,Accept, " +
			"X-Requested-With, Content-Type, " +
			"Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE")
		if r.Method == http.MethodOptions {
			return
		}
		handler.ServeHTTP(w, r)
	})
}
