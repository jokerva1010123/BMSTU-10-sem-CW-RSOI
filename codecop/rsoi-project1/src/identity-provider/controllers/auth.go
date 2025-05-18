package controllers

import (
	"encoding/json"
	"fmt"
	"identity-provider/models"
	"identity-provider/objects"
	"identity-provider/utils"
	"log"

	"net/http"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type auhtCtrl struct {
	auth *models.AuthM
}

func InitAuth(r *mux.Router, auth *models.AuthM) {
	ctrl := &auhtCtrl{auth}
	r.HandleFunc("/register", ctrl.register).Methods("POST")
	r.HandleFunc("/authorize", ctrl.authorize).Methods("POST")
}

type Token struct {
	jwt.StandardClaims
	Role string `json:"role,omitempty"`
}

const issuedAtLeewaySecs = 5

func (c *Token) Valid() error {
    c.StandardClaims.IssuedAt -= issuedAtLeewaySecs
    valid := c.StandardClaims.Valid()
    c.StandardClaims.IssuedAt += issuedAtLeewaySecs
    return valid
}

func newJWKs(rawJWKS string) *keyfunc.JWKS {
	jwksJSON := json.RawMessage(rawJWKS)
	jwks, err := keyfunc.NewJSON(jwksJSON)
	if err != nil {
		panic(err)
	}
	return jwks
}

func RetrieveToken(w http.ResponseWriter, r *http.Request) (*Token, error) {
	reqToken := r.Header.Get("Authorization")
	if len(reqToken) == 0 {
		TokenIsMissing(w)
		return nil, fmt.Errorf("TokenIsMissing")
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	tokenStr := splitToken[1]
	return processToken(tokenStr, w)
}

func processToken(tokenStr string, w http.ResponseWriter) (*Token, error) {
	jwks := newJWKs(utils.Config.RawJWKS)
	tk := &Token{}

	_, err := jwt.ParseWithClaims(tokenStr, tk, jwks.Keyfunc)
	if err != nil {
		JwtAccessDenied(w)
		return nil, fmt.Errorf("JwtAccessDenied: %s", err.Error())
	}
	if time.Now().Unix()-tk.ExpiresAt > 0 {
		TokenExpired(w)
		return nil, fmt.Errorf("TokenExpired")
	}

	return tk, nil
}

func (ctrl *auhtCtrl) register(w http.ResponseWriter, r *http.Request) {
	token, err := RetrieveToken(w, r)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if token.Role != utils.Admin.String() {
		Forbidden(w, fmt.Sprintf("not allowed for %s role", token.Role))
		return
	}

	req_body := new(objects.UserCreateRequest)
	err = json.NewDecoder(r.Body).Decode(req_body)
	if err != nil {
		log.Println(err.Error())
		ValidationErrorResponse(w, err.Error())
		return
	}

	err = ctrl.auth.Create(req_body)
	if err != nil {
		log.Println(err.Error())
		BadRequest(w, "user creation failed")
	} else {
		JsonSuccess(w, nil)
	}
}

func (ctrl *auhtCtrl) authorize(w http.ResponseWriter, r *http.Request) {
	req_body := new(objects.AuthRequest)
	err := json.NewDecoder(r.Body).Decode(req_body)
	if err != nil {
		log.Println(err.Error())
		ValidationErrorResponse(w, err.Error())
		return
	}

	data, err := ctrl.auth.Auth(req_body.Username, req_body.Password)
	if err != nil {
		log.Println(err.Error())
		BadRequest(w, "auth failed")
		return
	}

	token, err := processToken(data.AccessToken, w)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(token)
	data.Role = token.Role
	log.Println(data)

	JsonSuccess(w, data)
}
