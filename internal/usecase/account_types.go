package usecase

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountReqParams struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type GetAccountReqParams struct {
	Phone string `json:"phone"`
}

type CreateAccountResBody struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
