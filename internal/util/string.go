package util

import "unicode"

func UpperFirst(s string) string {
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}
