package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	var builder strings.Builder
	var bufRune rune
	var escapeMode bool

	for _, currentRune := range str {
		runeIsDigit := unicode.IsDigit(currentRune)
		runeIsSlash := currentRune == '\\'

		if bufRune != 0 {
			if runeIsDigit {
				counter, _ := strconv.Atoi(string(currentRune))
				builder.WriteString(strings.Repeat(string(bufRune), counter))
				bufRune = 0
				escapeMode = false
				continue
			} else {
				builder.WriteRune(bufRune)
			}
		}
		// first slash occurrence
		if !escapeMode && runeIsSlash {
			bufRune = 0
			escapeMode = true
			continue
		}
		// second slash occurrence or digit after slash or "not a digit" without slash
		if escapeMode && (runeIsDigit || runeIsSlash) || (!escapeMode && !runeIsDigit) {
			bufRune = currentRune
			escapeMode = false
		} else {
			return "", ErrInvalidString
		}
	}
	// If we have a rune without a counter at the end of the string
	if bufRune != 0 {
		builder.WriteRune(bufRune)
	}
	return builder.String(), nil
}
