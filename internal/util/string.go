package util

import "unicode"

func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	b := []rune(s)
	b[0] = unicode.ToUpper(b[0])
	return string(b)
}
