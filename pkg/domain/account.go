package domain

import "github.com/google/uuid"

type AccountID uuid.UUID

type Account struct {
	AccountID uuid.UUID
}
