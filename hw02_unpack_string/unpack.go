package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(original string) (string, error) {
	var result strings.Builder
	var previousChar rune

	if len(original) == 0 {
		return "", nil
	}

	for _, char := range original {
		repeatCount, convError := strconv.Atoi(string(char))

		if previousChar == 0 && convError == nil {
			return "", ErrInvalidString
		}

		if convError == nil && previousChar != 0 {
			result.WriteString(strings.Repeat(string(previousChar), repeatCount))
			previousChar = 0
			continue
		}

		if previousChar != 0 {
			result.WriteString(string(previousChar))
		}

		previousChar = char
	}

	if previousChar != 0 {
		result.WriteString(string(previousChar))
	}

	return result.String(), nil
}
