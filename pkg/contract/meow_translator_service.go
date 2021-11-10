package contract

type MeowTranslatorService interface {
	Translate(body string) string
}
