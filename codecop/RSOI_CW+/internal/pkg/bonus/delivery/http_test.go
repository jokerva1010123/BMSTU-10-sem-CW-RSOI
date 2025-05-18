package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/bonus"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetBonus(t *testing.T) {

	os.Setenv("SECRET", "ABCD")
	id := uuid.New()

	tokenModel := models.Token{
		UserUUID: id,
		UserRole: "auth",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	SecretKey := "ABCD"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenModel)

	jwtCookie, _ := token.SignedString([]byte(SecretKey))

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mock := bonus.NewMockIRepo(ctl)
	handler := NewBonusHandler(mock)

	mock.EXPECT().GetBonus(id).Return(0, models.StatusOkey).Times(1)
	r := httptest.NewRequest("GET", "/api/v1/miles", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"UUID": id.String(),
	})
	r.Header.Set("Authorization", "Bearer "+jwtCookie)
	w := httptest.NewRecorder()

	handler.GetBonus(w, r)
}

func TestSetBonus(t *testing.T) {
	os.Setenv("SECRET", "ABCD")
	id := uuid.New()

	tokenModel := models.Token{
		UserUUID: id,
		UserRole: "auth",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	SecretKey := "ABCD"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenModel)

	jwtCookie, _ := token.SignedString([]byte(SecretKey))

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mock := bonus.NewMockIRepo(ctl)
	handler := NewBonusHandler(mock)

	mock.EXPECT().SetBonus(id, 1).Return(1, models.StatusOkey).Times(1)
	r := httptest.NewRequest("GET", "/api/v1/miles", strings.NewReader(fmt.Sprint()))
	r = mux.SetURLVars(r, map[string]string{
		"UUID": id.String(),
	})
	r.Header.Set("Authorization", "Bearer "+jwtCookie)
	w := httptest.NewRecorder()

	handler.AddBonus(w, r)
}
