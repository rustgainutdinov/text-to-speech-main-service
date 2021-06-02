package domain

type TextToSpeechService interface {
	Translate(text string) (TranslationID, error)
}
