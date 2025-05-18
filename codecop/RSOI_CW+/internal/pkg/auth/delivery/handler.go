package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/auth"
	"RSOI_CW/internal/pkg/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/mailru/easyjson"
	"net/http"
	"os"
	"strings"
	"time"
)

type AuthHandler struct {
	repo auth.IRepo
}

func NewAuthHandler(repo auth.IRepo) *AuthHandler {
	return &AuthHandler{repo: repo}
}

func (h *AuthHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://3.67.182.34:8887")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method == http.MethodOptions {
		return
	}
	user := models.User{}
	err := easyjson.UnmarshalFromReader(r.Body, &user)
	if err != nil {
		middleware.Response(w, models.StatusBadUUID, nil)
		return
	}
	user, status := h.repo.AddUser(user)
	middleware.Response(w, status, user)
}

func (h *AuthHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://3.67.182.34:8887")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method == http.MethodOptions {
		return
	}
	users, status := h.repo.GetUsers()
	middleware.Response(w, status, users)
}

func (h *AuthHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	login, password, err := r.BasicAuth()
	if !err {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}
	user, state := h.repo.GetUser(login, password)
	if state != models.StatusOkey {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	tokenModel := models.Token{
		UserUUID: user.UUID,
		UserRole: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	SecretKey, err := os.LookupEnv("SECRET")
	if !err {
		panic("where is secret key")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenModel)

	jwtCookie, _ := token.SignedString([]byte(SecretKey))
	tokenResp := models.TokenResponse{jwtCookie}

	middleware.Response(w, models.StatusOkey, tokenResp)
}

func (h *AuthHandler) CheckToken(w http.ResponseWriter, r *http.Request) {
	var empty models.Token
	user := middleware.User(r)
	if user == empty {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}
	middleware.Response(w, models.StatusOkey, nil)

}

func (h *AuthHandler) CheckAdminToken(w http.ResponseWriter, r *http.Request) {

	var empty models.Token
	user := middleware.User(r)
	if user == empty {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}

	if strings.ToLower(user.UserRole) != "admin" {
		middleware.Response(w, models.StatusNoAuth, nil)
		return
	}
	middleware.Response(w, models.StatusOkey, nil)

}
