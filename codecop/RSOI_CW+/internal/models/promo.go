package models

import "github.com/google/uuid"

type Promo struct {
	PromoUUID uuid.UUID `json:"promo_uuid"`
	IsActive  bool      `json:"is_active"`
	Factor    float32   `json:"factor"`
}
