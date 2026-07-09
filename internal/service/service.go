package service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func GetFile(filePath string) (string, error) {
	return filepath.Abs(filePath)
}

func ReadFile(filePath string) (string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(fileData), nil
}

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

func Run(filePath string) (string, error) {
	absPath, err := GetFile(filePath)
	if err != nil {
		return "", fmt.Errorf("resolve path: %w", err)
	}

	content, err := ReadFile(absPath)
	if err != nil {
		return "", fmt.Errorf("read file: %w", err)
	}

	return DetectAndConvert(content)
}
