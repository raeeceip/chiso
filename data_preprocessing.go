package main

import (
	"strings"
	"unicode"
)

func preprocessText(text string) string {
	// Convert to lowercase
	text = strings.ToLower(text)

	// Remove extra whitespace
	text = strings.Join(strings.Fields(text), " ")

	// Remove punctuation
	text = strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, text)

	return text
}
