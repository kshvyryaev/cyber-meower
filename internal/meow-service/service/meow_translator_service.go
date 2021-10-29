package service

import (
	"regexp"
	"strings"

	"github.com/google/wire"
)

type MeowTranslatorService struct{}

func (service *MeowTranslatorService) Translate(body string) string {
	sb := strings.Builder{}
	digitRegexp, _ := regexp.Compile(`^\d+$`)
	letterRegexp, _ := regexp.Compile(`^[A-Za-zА-Яа-я]+$`)

	sb.WriteString("Translated to meow language:")

	for _, item := range body {
		isLetter := letterRegexp.MatchString(string(item))
		isDigit := digitRegexp.MatchString(string(item))

		if isLetter {
			sb.WriteString(" meow")
		} else if isDigit {
			sb.WriteString(" purr")
		} else {
			sb.WriteString(" sniff")
		}
	}
	sb.WriteString(" (translated from leather language: ")
	sb.WriteString(body)
	sb.WriteString(")")

	return sb.String()
}

func ProvideMeowTranslatorService() *MeowTranslatorService {
	return &MeowTranslatorService{}
}

var MeowTranslatorServiceSet = wire.NewSet(
	ProvideMeowTranslatorService,
)
