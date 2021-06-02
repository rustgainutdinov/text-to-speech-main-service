package domain

import "github.com/google/uuid"

type TranslationID uuid.UUID

type AccountTranslationRepo interface {
	Store(accountTranslation AccountTranslation) error
}

type AccountTranslation struct {
	AccountID
	TranslationID
}
