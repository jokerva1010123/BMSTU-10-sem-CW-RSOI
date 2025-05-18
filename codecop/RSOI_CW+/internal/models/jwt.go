package models

import "github.com/google/uuid"
import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserUUID           uuid.UUID //пользователь
	UserRole           string    //роль и права пользователя
	jwt.StandardClaims           //тут всё ясно
}

func (t Token) Valid() error {
	return nil
}

type TokenResponse struct {
	//Id    uuid.UUID
	Token string
}
