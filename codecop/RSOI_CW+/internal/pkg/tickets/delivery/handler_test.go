package delivery

import (
	"RSOI_CW/internal/models"
	"RSOI_CW/internal/pkg/tickets"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetMyTickets(t *testing.T) {

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

	repo := tickets.NewMockIRepo(ctl)
	handler := NewTicketHandler(repo)

	r := httptest.NewRequest("GET", "/api/v1/flights", strings.NewReader(fmt.Sprint()))
	w := httptest.NewRecorder()

	r.Header.Set("Authorization", "Bearer "+jwtCookie)

	repo.EXPECT().GetTickets(id).Return([]models.Ticket{}, 0).Times(1)
	handler.GetMyTickets(w, r)

}
