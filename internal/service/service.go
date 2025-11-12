package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(text string) bool {
	for _, letter := range text {
		if !strings.Contains(".- /", string(letter)) {
			return false
		}
	}
	return true
}

func ConvertAutomatically(input string) (string, error) {
	trimmed := strings.TrimSpace(input)

	if len(trimmed) == 0 {
		return "", fmt.Errorf("empty input")
	}

	if isMorse(trimmed) {
		return morse.ToText(trimmed), nil

	} else {
		return morse.ToMorse(trimmed), nil

	}
}
