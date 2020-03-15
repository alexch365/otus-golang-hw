package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	inputString = strings.TrimSpace(inputString)
	var builder strings.Builder
	var escapeMode bool

	for i := 0; i < len(inputString); i++ {
		lastChar := i == len(inputString)-1 //nolint:gomnd

		if string(inputString[i]) == `\` && !escapeMode {
			escapeMode = true
			if lastChar {
				return "", ErrInvalidString
			}
			continue
		}

		if !escapeMode && !unicode.IsLetter(rune(inputString[i])) {
			return "", ErrInvalidString
		}

		if !lastChar && unicode.IsDigit(rune(inputString[i+1])) {
			counter, _ := strconv.Atoi(string(inputString[i+1]))
			builder.WriteString(strings.Repeat(string(inputString[i]), counter))
			i++
		} else {
			builder.WriteString(string(inputString[i]))
		}
		escapeMode = false
	}
	return builder.String(), nil
}
