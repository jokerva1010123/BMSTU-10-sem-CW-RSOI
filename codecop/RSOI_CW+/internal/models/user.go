package models

import "github.com/google/uuid"


type Users []User

type User struct {
	Login    string    `json:"login"`
	Password string   `json:"-"`
	UUID     uuid.UUID `json:"id"`
	Role     string    `json:"role"`
}

