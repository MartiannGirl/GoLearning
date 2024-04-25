package main

import (
	"unicode/utf8"
)

func ToRunes(bytes []byte) []rune {
	result := []rune{}
	for i := 0; i < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])
		result = append(result, r)
		i += size
	}
	return result
}
