package middleware

import (
	"RSOI_CW/internal/models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
)

var DefaultError = models.Error{
	Message: "This is my error message",
}

var bearerPrefix = "Bearer "

func Response(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case models.StatusOkey:
		w.WriteHeader(http.StatusOK)
		if body != nil {
			jsn, err := json.Marshal(body)
			log.Print(err)
			_, _ = w.Write(jsn)
		}
	case models.StatusBadUUID:
		w.WriteHeader(http.StatusBadRequest)
		jsn, _ := json.Marshal(DefaultError)
		_, _ = w.Write(jsn)

	case models.StatusConflict:
		w.WriteHeader(http.StatusConflict)
		jsn, _ := json.Marshal(DefaultError)
		_, _ = w.Write(jsn)

	case models.StatusNotFound:
		w.WriteHeader(http.StatusNotFound)
		jsn, _ := json.Marshal(DefaultError)
		_, _ = w.Write(jsn)

	case models.StatusError:
		w.WriteHeader(http.StatusUnprocessableEntity)
		jsn, _ := json.Marshal(DefaultError)
		_, _ = w.Write(jsn)
	case models.StatusNoAuth:
		w.WriteHeader(http.StatusUnauthorized)
		jsn, _ := json.Marshal(DefaultError)
		_, _ = w.Write(jsn)
	default:
		w.WriteHeader(http.StatusTeapot)
	}

}

func UserUUID(r *http.Request) uuid.UUID {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return uuid.UUID{}
	}
	n := len(bearerPrefix)
	cookieValue := auth[n:]

	token, err := jwt.ParseWithClaims(cookieValue,
		&models.Token{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			SecretKey, _ := os.LookupEnv("SECRET")
			return []byte(SecretKey), nil
		})

	if err != nil {
		return uuid.UUID{}
	}

	tk := token.Claims.(*models.Token)
	return tk.UserUUID
}

func User(r *http.Request) models.Token {
	auth := r.Header.Get("Authorization")
	n := len(bearerPrefix)
	cookieValue := auth[n:]

	token, err := jwt.ParseWithClaims(cookieValue,
		&models.Token{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			SecretKey, _ := os.LookupEnv("SECRET")
			return []byte(SecretKey), nil
		})

	if err != nil {
		return models.Token{}
	}

	tk := token.Claims.(*models.Token)
	return *tk
}