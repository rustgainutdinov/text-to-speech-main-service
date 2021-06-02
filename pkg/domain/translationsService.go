package domain

type TranslationService interface {
	Translate(accountID AccountID, text string) (TranslationID, error)
}

type translationsService struct {
	textToSpeechService    TextToSpeechService
	accountTranslationRepo AccountTranslationRepo
}

func (t *translationsService) Translate(accountID AccountID, text string) (TranslationID, error) {
	translationID, err := t.textToSpeechService.Translate(text)
	if err != nil {
		return TranslationID{}, err
	}
	return translationID, t.accountTranslationRepo.Store(AccountTranslation{
		AccountID:     accountID,
		TranslationID: translationID,
	})
}
