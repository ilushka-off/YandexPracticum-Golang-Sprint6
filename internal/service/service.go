package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DetectAndConvert(str string) (string, error) {
	convertStr := strings.TrimSpace(str)

	if convertStr == "" {
		return "", errors.New("empty input")
	}

	isMorse := strings.IndexFunc(convertStr, func(r rune) bool {
		return !strings.ContainsRune(".- /", r)
	}) == -1

	if isMorse {
		return morse.ToText(convertStr), nil
	}

	return morse.ToMorse(convertStr), nil
}
