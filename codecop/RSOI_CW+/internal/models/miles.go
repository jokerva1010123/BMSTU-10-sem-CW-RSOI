package models

import "github.com/google/uuid"

type Bonus struct {
	UserUUID uuid.UUID `json:"user_uuid"`
	Balance  int       `json:"balance"`
}
