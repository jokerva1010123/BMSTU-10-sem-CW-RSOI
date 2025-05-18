package middleware

import (
	"io/ioutil"
	"net/http"
)

func CopyResponse(w http.ResponseWriter, resp *http.Response) {
	w.Header().Set("Content-Type", "application/json")
	if resp.StatusCode == http.StatusInternalServerError {
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		w.WriteHeader(resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err == nil || len(bodyBytes) != 0 {
		_, _ = w.Write(bodyBytes)
	}
}
